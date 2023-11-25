package main

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/user/service"
	"github.com/Yra-A/Fusion_Go/kitex_gen/user"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	klog.CtxDebugf(ctx, "UserRegister called: %s", req.GetUsername()+" "+req.GetPassword())
	resp = new(user.UserRegisterResponse)
	if req.Username == "" || req.Password == "" {
		resp.StatusCode = errno.EmptyUsernameOrPasswordErr.ErrCode
		resp.StatusMsg = errno.EmptyUsernameOrPasswordErr.ErrMsg
		return resp, nil
	}
	err = service.NewCreateUserService(ctx).CreateUser(req.Username, req.Password)
	if err == errno.UserAlreadyExistErr {
		resp.StatusCode = errno.UserAlreadyExistErr.ErrCode
		resp.StatusMsg = errno.UserAlreadyExistErr.ErrMsg
		return resp, nil
	}
	if err != nil {
		resp.StatusCode = errno.Fail.ErrCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, nil
	}
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	return resp, nil

}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	klog.CtxDebugf(ctx, "UserLogin called: %s", req.GetUsername()+" "+req.GetPassword())
	resp = new(user.UserLoginResponse)
	u, err := service.NewCheckUserService(ctx).CheckUser(req.Username, req.Password)
	if err != nil {
		resp.StatusCode = errno.InvalidCredentialsErr.ErrCode
		resp.StatusMsg = errno.InvalidCredentialsErr.ErrMsg
		return resp, nil
	}
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	resp.UserId = u
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	klog.CtxDebugf(ctx, "UserInfo called: %d", req.GetUserId())
	resp = new(user.UserInfoResponse)
	u, err := service.NewQueryUserService(ctx).QueryUser(req.UserId)
	if err != nil {
		resp.StatusCode = errno.Fail.ErrCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, err
	}
	if HasEmpty(u) {
		resp.StatusCode = errno.UserinfoNotSetErr.ErrCode
		resp.StatusMsg = errno.UserinfoNotSetErr.ErrMsg
		resp.UserInfo = u
		return resp, nil
	}
	resp.StatusCode = errno.SuccessCode
	resp.StatusMsg = errno.Success.ErrMsg
	resp.UserInfo = u
	return resp, nil
}
func HasEmpty(u *user.UserInfo) bool {
	return u.Gender == 0 || u.EnrollmentYear == 0 || u.MobilePhone == "" || u.College == "" || u.Nickname == "" || u.Realname == ""

}

// UserInfoUpload implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfoUpload(ctx context.Context, req *user.UserInfoUploadRequest) (resp *user.UserInfoUploadResponse, err error) {
	klog.CtxDebugf(ctx, "UserInfoUpload called: %d", req.UserInfo.UserId)
	resp = new(user.UserInfoUploadResponse)
	err = service.NewUploadUserService(ctx).UploadUserInfo(req.UserInfo)
	if err != nil {
		resp.StatusCode = errno.FailCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, nil
	}
	resp.StatusCode = errno.SuccessCode
	resp.StatusMsg = errno.Success.ErrMsg
	return
}

// UserProfileInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserProfileInfo(ctx context.Context, req *user.UserProfileInfoRequest) (resp *user.UserProfileInfoResponse, err error) {
	klog.CtxDebugf(ctx, "UserProfileInfo called: %d", req.GetUserId())
	resp = new(user.UserProfileInfoResponse)
	u, err := service.NewQueryUserProfileService(ctx).QueryUserProfile(req.UserId)
	if err != nil {
		resp.StatusCode = errno.FailCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, nil
	}
	resp.StatusCode = errno.SuccessCode
	resp.StatusMsg = errno.Success.ErrMsg
	resp.UserProfileInfo = u
	return resp, nil

}

// UserProfileUpload implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserProfileUpload(ctx context.Context, req *user.UserProfileUploadRequest) (resp *user.UserProfileUploadResponse, err error) {
	klog.CtxDebugf(ctx, "UserProfileUpload called: %d", req.GetUserId())
	resp = new(user.UserProfileUploadResponse)
	err = service.NewUploadUserService(ctx).UploadUserProfileInfo(req.UserProfileInfo)
	if err != nil {
		resp.StatusCode = errno.FailCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, nil
	}
	resp.StatusCode = errno.SuccessCode
	resp.StatusMsg = errno.Success.ErrMsg
	return
}
