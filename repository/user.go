package repository

import (
	"context"
	"gin-web/models"
	"gin-web/repository/cache"
	"gin-web/repository/dao"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	dao   *dao.UserDAO
	cache *cache.UserCache
}

var userRepository *UserRepository

func NewUserRepository(ctx *gin.Context) *UserRepository {
	if userRepository == nil {
		userRepository = &UserRepository{
			dao:   dao.NewUserDAO(ctx),
			cache: cache.NewUserCache(ctx),
		}
	}
	return userRepository
}

func toModelUser(u dao.User) models.User {
	return models.User{
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
	return u.dao.Insert(ctx, dao.User{
		Email:    user.Email,
		Password: *user.Password,
	})
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (models.User, error) {
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
	return u.dao.AssociateRoles(ctx, uid, roles)
}
