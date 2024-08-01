package router

import (
	v1 "gin-web/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRoleRouter(r *gin.RouterGroup) {
	roleApi := v1.NewRoleApi()
	roleAccessGroup := r.Group("role")
	{
		// TODO:添加权限限制
		roleAccessGroup.POST("create", roleApi.CreateRole)
	}
}
