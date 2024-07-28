// Code generated by hertz generator.

package admin

import (
	"context"
	"formulago/biz/domain"
	logic "formulago/biz/logic/admin"
	"formulago/data"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	"strconv"

	"formulago/api/model/admin"
	base "formulago/api/model/base"
	"github.com/cloudwego/hertz/pkg/app"
)

// CreateAuthority .
// @router /api/admin/authority/api/create [POST]
func CreateAuthority(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.CreateOrUpdateApiAuthorityReq
	resp := new(base.BaseResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = base.ErrCode_Fail
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	var apiInfos []*domain.ApiAuthorityInfo
	err = copier.Copy(&apiInfos, &req.Data)
	if err != nil {
		resp.ErrCode = base.ErrCode_Fail
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	err = logic.NewAuthority(data.Default(), data.CasbinEnforcer()).UpdateApiAuthority(ctx, strconv.Itoa(int(req.RoleID)), apiInfos)
	if err != nil {
		resp.ErrCode = base.ErrCode_Fail
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// UpdateApiAuthority .
// @router /api/admin/authority/api/update [POST]
func UpdateApiAuthority(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.CreateOrUpdateApiAuthorityReq
	resp := new(base.BaseResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = base.ErrCode_Fail
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	var apiInfos []*domain.ApiAuthorityInfo
	err = copier.Copy(&apiInfos, &req.Data)
	if err != nil {
		resp.ErrCode = base.ErrCode_Fail
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	err = logic.NewAuthority(data.Default(), data.CasbinEnforcer()).UpdateApiAuthority(ctx, strconv.Itoa(int(req.RoleID)), apiInfos)
	if err != nil {
		resp.ErrCode = base.ErrCode_Fail
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// ApiAuthority .
// @router /api/admin/authority/api/role [POST]
// ApiAuthority Get Api authority policy by role id
func ApiAuthority(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	resp := new(admin.ApiAuthorityListInfoResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = base.ErrCode_Fail
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	policies, err := logic.NewAuthority(data.Default(), data.CasbinEnforcer()).ApiAuthority(ctx, strconv.Itoa(int(req.ID)))
	if err != nil {
		resp.ErrCode = base.ErrCode_Fail
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.Total = uint64(len(policies))
	for _, v := range policies {
		resp.Data = append(resp.Data, &admin.ApiAuthorityInfo{
			Path:   v.Path,
			Method: v.Method,
		})
	}
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// CreateMenuAuthority .
// @router /api/admin/authority/menu/create [POST]
func CreateMenuAuthority(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.MenuAuthorityInfoReq
	resp := new(base.BaseResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = base.ErrCode_Fail
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	err = logic.NewAuthority(data.Default(), data.CasbinEnforcer()).UpdateMenuAuthority(ctx, req.RoleID, req.MenuIDs)
	if err != nil {
		resp.ErrCode = base.ErrCode_Fail
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"

	c.JSON(consts.StatusOK, resp)
}

// UpdateMenuAuthority .
// @router /api/admin/authority/menu/update [POST]
func UpdateMenuAuthority(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.MenuAuthorityInfoReq
	resp := new(base.BaseResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = base.ErrCode_Fail
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	err = logic.NewAuthority(data.Default(), data.CasbinEnforcer()).UpdateMenuAuthority(ctx, req.RoleID, req.MenuIDs)
	if err != nil {
		resp.ErrCode = base.ErrCode_Fail
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"

	c.JSON(consts.StatusOK, resp)
}

// MenuAuthority .
// @router /api/admin/authority/menu/role [POST]
func MenuAuthority(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	resp := new(admin.MenuAuthorityInfoResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = base.ErrCode_Fail
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	menuIDs, err := logic.NewAuthority(data.Default(), data.CasbinEnforcer()).MenuAuthority(ctx, req.ID)
	if err != nil {
		resp.ErrCode = base.ErrCode_Fail
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	resp.RoleID = req.ID
	resp.MenuIDs = menuIDs
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"

	c.JSON(consts.StatusOK, resp)
}