package service

import (
	"context"
	"database/sql"
	"errors"
	"gin-web/pkg/database"
	"gin-web/pkg/global"
	"gin-web/pkg/redis"
	"go.uber.org/zap"
	"runtime/debug"
	"sync"
)

type BasicService struct {
	logger *zap.SugaredLogger
	redis  *redis.CommonRedisClient
}

var (
	basicOnce sync.Once
	basic     *BasicService
)

func NewBasicService() *BasicService {
	basicOnce.Do(func() {
		basic = &BasicService{
			logger: global.Logger,
			redis:  global.RedisClient,
		}
	})
	return basic
}

// Transaction 开启事务,join为true则加入上下文中的事务,如果上下文中没有事务则开启新事务,join为false时直接开启新事务
func (s *BasicService) Transaction(ctx context.Context, join bool, fn database.Action, options ...*sql.TxOptions) error {
	isStarter := false // 是否是事务开启者
	manager := &database.TransactionManager{
		DB:                           global.DB,
		AlreadyCommittedOrRolledBack: false,
	}
	if join {
		if m := database.LoadManager(ctx); m != nil {
			// 从上下文中查找到已经存在的事务
			manager = m
		} else {
			// 未找到已经存在的事务则开启新事务
			isStarter = true
			manager.DB = manager.DB.Begin(options...).WithContext(ctx)
		}
	} else {
		// 开启新事务
		isStarter = true
		manager.DB = manager.DB.Begin(options...).WithContext(ctx)
	}
	if manager.DB.Error != nil {
		return manager.DB.Error
	}
	// 将TransactionManager放入上下文
	wrapContext := database.InjectManager(ctx, manager)
	var execErr error
	wrapFunc := func() {
		defer func() {
			if err := recover(); err != nil {
				stackInfo := string(debug.Stack())
				s.logger.Errorf("Transaction panic,堆栈信息:%s", stackInfo)
				execErr = errors.New(stackInfo)
			}
		}()
		execErr = fn(wrapContext)
	}
	wrapFunc()
	if execErr != nil {
		if !manager.AlreadyCommittedOrRolledBack {
			manager.AlreadyCommittedOrRolledBack = true
			if rollback := manager.DB.Rollback(); rollback.Error != nil {
				return rollback.Error
			}
		}
		return execErr
	}
	if isStarter && !manager.AlreadyCommittedOrRolledBack {
		manager.AlreadyCommittedOrRolledBack = true
		if commit := manager.DB.Commit(); commit.Error != nil {
			return commit.Error
		}
	}
	return nil
}
