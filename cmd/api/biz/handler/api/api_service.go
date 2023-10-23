// Code generated by hertz generator.

package api

import (
	"Fusion/cmd/api/biz/handler"
	api "Fusion/cmd/api/biz/model/api"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

// UserRegister .
// @router /fusion/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UserRegisterResponse)

	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router /fusion/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UserLoginResponse)

	c.JSON(consts.StatusOK, resp)
}

// UserInfo .
// @router /fusion/user/info/ [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserInfoRequest
	if err = c.BindAndValidate(&req); err != nil {
		handler.BadResponse(c, err)
		return
	}

	resp := new(api.UserInfoResponse)

	c.JSON(consts.StatusOK, resp)
}

// Todo: 未完成
// UserProfileInfo .
// @router /fusion/user/profile/{user_id} [GET]
func UserProfileInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserProfileInfoRequest
	if err = c.BindAndValidate(&req); err != nil {
		handler.BadResponse(c, err)
		return
	}
	if user_id := c.Param("user_id"); user_id != "" {
		tempID, err := strconv.Atoi(user_id)
		if err != nil {
			handler.BadResponse(c, err)
			return
		}
		req.UserID = int32(tempID)
	}
	if token := c.Query("token"); token != "" {
		req.Token = token
	}

	resp := new(api.UserProfileInfoResponse)

	c.JSON(consts.StatusOK, resp)
}

// UserProfileUpload .
// @router /fusion/user/profile/upload/ [POST]
func UserProfileUpload(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserProfileUploadRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UserProfileUploadResponse)

	c.JSON(consts.StatusOK, resp)
}

// ContestList .
// @router /fusion/contest/list/ [GET]
func ContestList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ContestListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.ContestListResponse)

	c.JSON(consts.StatusOK, resp)
}

// ContestInfo .
// @router /fusion/contest/info/{contest_id} [GET]
func ContestInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ContestInfoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.ContestInfoResponse)

	c.JSON(consts.StatusOK, resp)
}
