package repository

import (
	"context"
	"gin-web/models"
	"gin-web/repository/cache"
	"gin-web/repository/dao"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	dao   *dao.UserDAO
	cache *cache.UserCache
}

func NewUserRepository(ctx *gin.Context) *UserRepository {
	return &UserRepository{
		dao:   dao.NewUserDAO(ctx),
		cache: cache.NewUserCache(ctx),
	}
}

func toModelUser(u *dao.User) *models.User {
	return &models.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: &u.Password,
		NickName: u.NickName,
		Gender:   models.UserGender(u.Gender),
		About:    u.About,
		Birthday: time.UnixMilli(u.Birthday).Format(time.DateTime),
	}
}

func (u *UserRepository) Create(ctx context.Context, user models.User) error {
	return u.dao.Insert(ctx, &dao.User{
		Email:    user.Email,
		Password: *user.Password,
	})
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := u.dao.FindByEmail(ctx, email)
	return toModelUser(user), err
}

func (u *UserRepository) CacheTokenPair(ctx context.Context, email string, pair *models.TokenPair) error {
	return u.cache.HSetTokenPair(ctx, email, pair)
}

func (u *UserRepository) TokenPairExist(ctx context.Context, email string) (bool, error) {
	return u.cache.HExistsTokenPair(ctx, email)
}

func (u *UserRepository) DelTokenPair(ctx context.Context, email string) error {
	return u.cache.HDelTokenPair(ctx, email)
}

func (u *UserRepository) AssociateRoles(ctx context.Context, uid uint, role_ids []uint) error {
	var roles []dao.Role
	for _, id := range role_ids {
		roles = append(roles, dao.Role{
			Model: gorm.Model{
				ID: id,
			},
		})
	}
	return u.dao.AssociateRoles(ctx, uid, &roles)
}

func (u *UserRepository) FindByUid(ctx context.Context, uid uint) (*models.User, error) {
	user, err := u.dao.FindByUid(ctx, uid)
	return toModelUser(user), err
}

func (u *UserRepository) FindRolesByUid(ctx context.Context, uid uint) ([]*models.RoleWithoutUsers, error) {
	roles, err := u.dao.FindRolesByUid(ctx, uid)
	if err != nil {
		return nil, err
	}
	result := lo.Map[*dao.PureRole, *models.RoleWithoutUsers](roles, func(item *dao.PureRole, _ int) *models.RoleWithoutUsers {
		return toModelRole(item)
	})
	return result, nil
}
