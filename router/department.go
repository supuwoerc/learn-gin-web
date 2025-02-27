package router

import (
	v1 "gin-web/api/v1"
	"github.com/gin-gonic/gin"
)

func InitDepartmentApiRouter(r *gin.RouterGroup) {
	departmentApi := v1.NewDepartmentApi()
	departmentAccessGroup := r.Group("department")
	{
		departmentAccessGroup.POST("create", departmentApi.CreateDepartment)
		departmentAccessGroup.GET("tree", departmentApi.GetDepartmentTree)
	}
}
