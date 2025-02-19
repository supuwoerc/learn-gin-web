package v1

import (
	"context"
	"gin-web/pkg/request"
	"gin-web/pkg/response"
	"gin-web/pkg/utils"
	"gin-web/service"
	"github.com/gin-gonic/gin"
	"sync"
)

type PermissionService interface {
	CreatePermission(ctx context.Context, operator uint, name, resource string, roleIds []uint) error
	GetPermissionList(ctx context.Context, keyword string, limit, offset int) ([]*response.PermissionListRowResponse, int64, error)
	GetPermissionDetail(ctx context.Context, id uint) (*response.PermissionDetailResponse, error)
	UpdatePermission(ctx context.Context, operator uint, id uint, name, resource string, roleIds []uint) error
	DeletePermission(ctx context.Context, id, operator uint) error
}

type PermissionApi struct {
	*BasicApi
	service PermissionService
}

var (
	permissionOnce sync.Once
	permissionApi  *PermissionApi
)

func NewPermissionApi() *PermissionApi {
	permissionOnce.Do(func() {
		permissionApi = &PermissionApi{
			BasicApi: NewBasicApi(),
			service:  service.NewPermissionService(),
		}
	})
	return permissionApi
}

func (r *PermissionApi) CreatePermission(ctx *gin.Context) {
	var params request.CreatePermissionRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ParamsValidateFail(ctx, err)
		return
	}
	claims, err := utils.GetContextClaims(ctx)
	if err != nil || claims == nil {
		response.FailWithCode(ctx, response.AuthErr)
		return
	}
	err = r.service.CreatePermission(ctx, claims.User.ID, params.Name, params.Resource, params.Roles)
	if err != nil {
		response.FailWithError(ctx, err)
		return
	}
	response.Success(ctx)
}

func (r *PermissionApi) GetPermissionList(ctx *gin.Context) {
	var params request.GetPermissionListRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.ParamsValidateFail(ctx, err)
		return
	}
	list, total, err := r.service.GetPermissionList(ctx, params.Keyword, params.Limit, params.Offset)
	if err != nil {
		response.FailWithError(ctx, err)
		return
	}
	response.SuccessWithPageData(ctx, total, list)
}

func (r *PermissionApi) GetPermissionDetail(ctx *gin.Context) {
	var params request.GetPermissionDetailRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.ParamsValidateFail(ctx, err)
		return
	}
	detail, err := r.service.GetPermissionDetail(ctx, params.ID)
	if err != nil {
		response.FailWithError(ctx, err)
		return
	}
	response.SuccessWithData(ctx, detail)
}

func (r *PermissionApi) UpdatePermission(ctx *gin.Context) {
	var params request.UpdatePermissionRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ParamsValidateFail(ctx, err)
		return
	}
	claims, err := utils.GetContextClaims(ctx)
	if err != nil || claims == nil {
		response.FailWithCode(ctx, response.AuthErr)
		return
	}
	err = r.service.UpdatePermission(ctx, claims.User.ID, params.ID, params.Name, params.Resource, params.Roles)
	if err != nil {
		response.FailWithError(ctx, err)
		return
	}
	response.Success(ctx)
}

func (r *PermissionApi) DeletePermission(ctx *gin.Context) {
	var params request.DeletePermissionRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ParamsValidateFail(ctx, err)
		return
	}
	claims, err := utils.GetContextClaims(ctx)
	if err != nil || claims == nil {
		response.FailWithCode(ctx, response.AuthErr)
		return
	}
	err = r.service.DeletePermission(ctx, params.ID, claims.User.ID)
	if err != nil {
		response.FailWithError(ctx, err)
		return
	}
	response.Success(ctx)
}
