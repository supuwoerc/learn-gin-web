// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bootstrap

import (
	"gin-web/api/v1"
	"gin-web/initialize"
	"gin-web/middleware"
	"gin-web/pkg/captcha"
	"gin-web/pkg/email"
	"gin-web/pkg/job"
	"gin-web/pkg/jwt"
	"gin-web/pkg/utils"
	"gin-web/repository"
	"gin-web/repository/cache"
	"gin-web/repository/dao"
	"gin-web/router"
	"gin-web/service"
)

// Injectors from wire.go:

func WireApp() *App {
	viper := initialize.NewViper()
	writeSyncer := initialize.NewWriterSyncer(viper)
	sugaredLogger := initialize.NewZapLogger(viper, writeSyncer)
	cronLogger := initialize.NewCronLogger(sugaredLogger)
	cron := initialize.NewCronClient(cronLogger)
	jobRegisterer := job.NewJobRegisterer(cronLogger, cron, sugaredLogger, viper)
	dialer := initialize.NewDialer(viper)
	emailClient := email.NewEmailClient(sugaredLogger, dialer, viper)
	engine := initialize.NewEngine(writeSyncer, emailClient, sugaredLogger, viper)
	httpServer := initialize.NewServer(viper, engine, sugaredLogger)
	routerGroup := router.NewRouter(engine, viper)
	db := initialize.NewGORM(viper)
	commonRedisClient := initialize.NewRedisClient(writeSyncer, viper)
	redisLocksmith := utils.NewRedisLocksmith(sugaredLogger, commonRedisClient)
	basicService := service.NewBasicService(sugaredLogger, db, redisLocksmith, viper)
	basicDAO := dao.NewBasicDao(db)
	attachmentDAO := dao.NewAttachmentDAO(basicDAO)
	attachmentRepository := repository.NewAttachmentRepository(attachmentDAO)
	attachmentService := service.NewAttachmentService(basicService, attachmentRepository)
	userDAO := dao.NewUserDAO(basicDAO)
	userCache := cache.NewUserCache(commonRedisClient)
	userRepository := repository.NewUserRepository(userDAO, userCache)
	tokenBuilder := jwt.NewJwtBuilder(db, commonRedisClient, viper, userRepository)
	authMiddleware := middleware.NewAuthMiddleware(viper, userRepository, tokenBuilder)
	attachmentApi := v1.NewAttachmentApi(routerGroup, attachmentService, authMiddleware, viper)
	redisStore := captcha.NewRedisStore(commonRedisClient, viper)
	captchaService := service.NewCaptchaService(redisStore)
	captchaApi := v1.NewCaptchaApi(routerGroup, captchaService)
	departmentDAO := dao.NewDepartmentDAO(basicDAO)
	departmentCache := cache.NewDepartmentCache(commonRedisClient)
	departmentRepository := repository.NewDepartmentRepository(departmentDAO, departmentCache)
	departmentService := service.NewDepartmentService(basicService, departmentRepository, userRepository)
	departmentApi := v1.NewDepartmentApi(routerGroup, departmentService, authMiddleware)
	permissionDAO := dao.NewPermissionDAO(basicDAO)
	permissionRepository := repository.NewPermissionRepository(permissionDAO)
	roleDAO := dao.NewRoleDAO(basicDAO)
	roleRepository := repository.NewRoleRepository(roleDAO)
	permissionService := service.NewPermissionService(basicService, permissionRepository, roleRepository)
	permissionApi := v1.NewPermissionApi(routerGroup, permissionService, authMiddleware)
	pingService := service.NewPingService(basicService)
	pingApi := v1.NewPingApi(routerGroup, pingService, authMiddleware)
	roleService := service.NewRoleService(basicService, roleRepository, userRepository, permissionRepository)
	roleApi := v1.NewRoleApi(routerGroup, roleService, authMiddleware)
	userService := service.NewUserService(basicService, captchaService, roleRepository, userRepository, emailClient, tokenBuilder)
	userApi := v1.NewUserApi(routerGroup, userService, authMiddleware)
	app := &App{
		logger:        sugaredLogger,
		viper:         viper,
		jobRegisterer: jobRegisterer,
		httpServer:    httpServer,
		attachmentApi: attachmentApi,
		captchaApi:    captchaApi,
		departmentApi: departmentApi,
		permissionApi: permissionApi,
		pingApi:       pingApi,
		roleApi:       roleApi,
		userApi:       userApi,
	}
	return app
}
