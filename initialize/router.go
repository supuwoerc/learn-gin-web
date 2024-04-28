package initialize

import (
	"gin-web/middleware"
	"gin-web/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 不携带日志和Recovery中间件，自己添加中间件，为了方便收集Recovery日志
	r := gin.New()
	// 控制台logger中间件
	r.Use(gin.Logger())
	// recovery中间件
	r.Use(middleware.Recovery())
	// 跨域中间件
	r.Use(middleware.Cors())
	group := r.Group("api")
	router.InitPingRouter(group)
	router.InitPublicRouter(group)
	return r
}
