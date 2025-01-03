package bootstrap

import (
	"gin-web/initialize"
	"gin-web/pkg/global"
)

func Start() {
	initialize.InitConfig()
	global.Logger = initialize.InitZapLogger()
	global.DB = initialize.InitGORM()
	global.RedisClient = initialize.InitRedis()
	global.Localizer = initialize.InitI18N()
	initialize.InitServer(initialize.InitEngine(initialize.LoggerSyncer))
}
