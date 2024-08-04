// Code generated by hertz generator.

package admin

import (
	"context"
	"github.com/jinzhu/copier"
	"hertz-admin/biz/domain"
	"hertz-admin/biz/handler/middleware"
	logic "hertz-admin/biz/logic/admin"
	"hertz-admin/configs"
	"hertz-admin/data"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	admin "hertz-admin/api/model/admin"
	base "hertz-admin/api/model/base"
)

// CreateProvider .
// @router /api/admin/oauth/provider/create [POST]
func CreateProvider(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.ProviderInfo
	resp := new(base.BaseResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	var providerInfo domain.ProviderInfo
	err = copier.Copy(&providerInfo, &req)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	err = logic.NewOauth(data.Default(), configs.Data()).Create(ctx, &providerInfo)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.StatusCode = base.StatusCode_Success
	resp.StatusMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// UpdateProvider .
// @router /api/admin/oauth/provider/update [POST]
func UpdateProvider(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.ProviderInfo
	resp := new(base.BaseResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	var providerInfo domain.ProviderInfo
	err = copier.Copy(&providerInfo, &req)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	err = logic.NewOauth(data.Default(), configs.Data()).Update(ctx, &providerInfo)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.StatusCode = base.StatusCode_Success
	resp.StatusMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// DeleteProvider .
// @router /api/admin/oauth/provider [DELETE]
func DeleteProvider(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	resp := new(base.BaseResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		return
	}

	err = logic.NewOauth(data.Default(), configs.Data()).Delete(ctx, req.ID)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		return
	}

	resp.StatusCode = base.StatusCode_Success
	resp.StatusMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// GetProviderList .
// @router /api/admin/oauth/provider/list [GET]
func GetProviderList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.ProviderListReq
	resp := new(admin.ProviderListResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	var ListReq domain.OauthListReq
	ListReq.Page = req.Page
	ListReq.PageSize = req.PageSize
	ListReq.Name = req.Name
	l, total, err := logic.NewOauth(data.Default(), configs.Data()).List(ctx, &ListReq)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	resp.StatusCode = base.StatusCode_Success
	resp.StatusMsg = "success"
	resp.Data = &admin.ProviderInfoList{
		Total: uint64(total),
		Data:  make([]*admin.ProviderInfo, 0, len(l)),
	}
	err = copier.Copy(&resp.Data.Data, &l)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// OauthLogin .
// @router /api/oauth/login [POST]
func OauthLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.OauthLoginReq
	resp := new(admin.OauthRedirectResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		return
	}

	var loginReq domain.OauthLoginReq
	loginReq.Provider = req.Provider
	loginReq.State = req.State
	url, err := logic.NewOauth(data.Default(), configs.Data()).Login(ctx, &loginReq)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		return
	}

	resp.Data = &admin.OauthRedirectInfo{
		Url: url,
	}
	resp.StatusCode = base.StatusCode_Success
	resp.StatusMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// OauthCallback .
// @router /api/oauth/callback [GET]
func OauthCallback(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.CallbackReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, err.Error())
		return
	}

	stateList := strings.Split(req.State, "-")
	if len(stateList) < 2 {
		c.JSON(consts.StatusBadRequest, "state error, format should be: {random string}-{provider}")
		return
	}
	providerName := stateList[1]
	var callbackReq domain.OauthCallbackReq
	callbackReq.ProviderName = providerName
	callbackReq.Code = req.Code
	callbackReq.State = req.State
	userInfo, err := logic.NewOauth(data.Default(), configs.Data()).Callback(ctx, &callbackReq)
	if err != nil {
		c.JSON(consts.StatusBadRequest, err.Error())
		return
	}

	// goto jwt login middleware handler
	ctx = context.WithValue(ctx, "OAuthKey", configs.Data().Auth.OAuthKey)
	c.Set("provider", callbackReq.ProviderName)
	c.Set("credential", userInfo.Credential)
	middleware.GetJWTMiddleware(configs.Data(), data.Default(), data.CasbinEnforcer()).LoginHandler(ctx, c)
}
