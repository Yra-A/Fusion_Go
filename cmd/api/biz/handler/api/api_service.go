// Code generated by hertz generator.

package api

import (
    "context"
    "fmt"
    "github.com/Yra-A/Fusion_Go/cmd/api/biz/handler"
    "github.com/Yra-A/Fusion_Go/cmd/api/biz/model/api"
    "github.com/Yra-A/Fusion_Go/cmd/api/biz/mw/jwt"
    "github.com/Yra-A/Fusion_Go/cmd/api/biz/mw/oss"
    "github.com/Yra-A/Fusion_Go/cmd/api/rpc"
    "github.com/Yra-A/Fusion_Go/kitex_gen/user"
    conf "github.com/Yra-A/Fusion_Go/pkg/configs/oss"
    "github.com/Yra-A/Fusion_Go/pkg/errno"
    "github.com/Yra-A/Fusion_Go/pkg/utils"
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/common/hlog"
    "github.com/cloudwego/hertz/pkg/protocol/consts"
    "io"
    "strconv"
    "time"
)

// UserRegister .
// @router /fusion/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRegisterRequest
	if err = c.BindAndValidate(&req); err != nil {
		handler.BadResponse(c, err)
		return
	}
	kresp, err := rpc.UserRegister(context.Background(), &user.UserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	resp := new(api.UserRegisterResponse)
	resp.StatusCode = kresp.StatusCode
	resp.StatusMsg = kresp.StatusMsg
	handler.SendResponse(c, resp)
}

// UserLogin .
// @router /fusion/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
    jwt.JwtMiddleware.LoginHandler(ctx, c)
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

    kresp, err := rpc.UserInfo(context.Background(), &user.UserInfoRequest{
        UserId: req.UserID,
    })
    if err != nil {
        handler.BadResponse(c, err)
        return
    }
    u := kresp.UserInfo
    resp := new(api.UserInfoResponse)
    resp.StatusCode = errno.Success.ErrCode
    resp.StatusMsg = errno.Success.ErrMsg
    resp.UserInfo = &api.UserInfo{
        UserID:         u.UserId,
        Gender:         u.Gender,
        EnrollmentYear: u.EnrollmentYear,
        MobilePhone:    u.MobilePhone,
        College:        u.College,
        Nickname:       u.Nickname,
        Realname:       u.Realname,
        HasProfile:     u.HasProfile,
        AvatarURL:      u.AvatarUrl,
    }
    handler.SendResponse(c, resp)

}

// UserInfoUpload .
// @router /fusion/user/info/upload/ [POST]
func UserInfoUpload(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserInfoUploadRequest
	if err = c.BindAndValidate(&req); err != nil {
		handler.BadResponse(c, err)
		return
	}
	kresp, err := rpc.UserInfoUpload(context.Background(), &user.UserInfoUploadRequest{
		UserInfo: &user.UserInfo{
			UserId:         req.UserInfo.UserID,
			Gender:         req.UserInfo.Gender,
			EnrollmentYear: req.UserInfo.EnrollmentYear,
			MobilePhone:    req.UserInfo.MobilePhone,
			College:        req.UserInfo.College,
			Nickname:       req.UserInfo.Nickname,
			Realname:       req.UserInfo.Realname,
			HasProfile:     req.UserInfo.HasProfile,
			AvatarUrl:      req.UserInfo.AvatarURL,
		},
	})
	if err != nil {
		handler.BadResponse(c, err)
		return
	}
	resp := new(api.UserInfoUploadResponse)
	resp.StatusCode = kresp.StatusCode
	resp.StatusMsg = kresp.StatusMsg
	handler.SendResponse(c, resp)
}

// UserProfileInfo .
// @router /fusion/user/profile/{user_id} [GET]
func UserProfileInfo(ctx context.Context, c *app.RequestContext) {
    var err error
    var req api.UserProfileInfoRequest
    if err = c.BindAndValidate(&req); err != nil {
        handler.BadResponse(c, err)
        return
    }

    kresp, err := rpc.UserProfileInfo(context.Background(), &user.UserProfileInfoRequest{
        UserId: req.UserID,
    })

    if err != nil {
        handler.BadResponse(c, err)
        return
    }

	u := kresp.UserProfileInfo
	resp := new(api.UserProfileInfoResponse)
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	resp.UserProfileInfo = &api.UserProfileInfo{
		Introduction: u.Introduction,
		QqNumber:     u.QqNumber,
		WechatNumber: u.WechatNumber,
		Honors:       u.Honors,
		Images:       u.Images,
		UserInfo:     utils.ConvertUserToAPI(u.UserInfo),
	}
	handler.SendResponse(c, resp)
}

// UserProfileUpload .
// @router /fusion/user/profile/upload/ [POST]
func UserProfileUpload(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserProfileUploadRequest
	if err = c.BindAndValidate(&req); err != nil {
		handler.BadResponse(c, err)
		return
	}
	if req.GetUserID() != req.UserProfileInfo.UserInfo.UserID {
		handler.BadResponse(c, errno.ParamErr)
		return
	}
	kresp, err := rpc.UserProfileUpload(context.Background(), &user.UserProfileUploadRequest{
		UserId: req.UserID,
		UserProfileInfo: &user.UserProfileInfo{
			Introduction: req.UserProfileInfo.Introduction,
			QqNumber:     req.UserProfileInfo.QqNumber,
			WechatNumber: req.UserProfileInfo.WechatNumber,
			Honors:       req.UserProfileInfo.Honors,
			Images:       req.UserProfileInfo.Images,
			UserInfo:     utils.ConvertAPIToUser(req.UserProfileInfo.UserInfo),
		},
	})
	if err != nil {
		handler.BadResponse(c, err)
		return
	}
	resp := new(api.UserProfileUploadResponse)
	resp.StatusCode = kresp.StatusCode
	resp.StatusMsg = kresp.StatusMsg
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

// ImageUpload .
// @router /fusion/utils/upload/img [POST]
func ImageUpload(ctx context.Context, c *app.RequestContext) {
    var err error
    var req api.ImageUploadRequest
    err = c.BindAndValidate(&req)
    if err != nil {
        handler.BadResponse(c, err)
        return
    }
    file, err := c.FormFile("file")
    if err != nil {
        handler.BadResponse(c, err)
        return
    }
    src, err := file.Open()
    defer src.Close()
    if err != nil {
        handler.BadResponse(c, err)
        return
    }

    bytes, err := io.ReadAll(src)
    if err != nil {
        handler.BadResponse(c, err)
        return
    }
    req.File = bytes

    imageName := utils.NewImageName(time.Now().Unix())

    err = oss.UploadFile(imageName, req.File)
    hlog.CtxInfof(ctx, "图片上传大小为:"+strconv.FormatInt(int64(len(req.File)), 10)+"B")
    if err != nil {
        hlog.CtxInfof(ctx, "上传图片出现错误: "+err.Error())
    }

    imgURL := fmt.Sprintf("%s%s", conf.PublicURL, imageName)

    resp := new(api.ImageUploadResponse)
    resp.StatusCode = 0
    resp.StatusMsg = "success upload image!"
    resp.ImageURL = imgURL

    handler.SendResponse(c, resp)
}
