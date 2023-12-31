// Code generated by Kitex v0.7.3. DO NOT EDIT.

package userservice

import (
	"context"
	user "github.com/Yra-A/Fusion_Go/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"UserRegister":      kitex.NewMethodInfo(userRegisterHandler, newUserServiceUserRegisterArgs, newUserServiceUserRegisterResult, false),
		"UserLogin":         kitex.NewMethodInfo(userLoginHandler, newUserServiceUserLoginArgs, newUserServiceUserLoginResult, false),
		"UserInfo":          kitex.NewMethodInfo(userInfoHandler, newUserServiceUserInfoArgs, newUserServiceUserInfoResult, false),
		"UserInfoUpload":    kitex.NewMethodInfo(userInfoUploadHandler, newUserServiceUserInfoUploadArgs, newUserServiceUserInfoUploadResult, false),
		"UserProfileInfo":   kitex.NewMethodInfo(userProfileInfoHandler, newUserServiceUserProfileInfoArgs, newUserServiceUserProfileInfoResult, false),
		"UserProfileUpload": kitex.NewMethodInfo(userProfileUploadHandler, newUserServiceUserProfileUploadArgs, newUserServiceUserProfileUploadResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "user",
		"ServiceFilePath": `idl/user.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.3",
		Extra:           extra,
	}
	return svcInfo
}

func userRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserRegisterArgs)
	realResult := result.(*user.UserServiceUserRegisterResult)
	success, err := handler.(user.UserService).UserRegister(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserRegisterArgs() interface{} {
	return user.NewUserServiceUserRegisterArgs()
}

func newUserServiceUserRegisterResult() interface{} {
	return user.NewUserServiceUserRegisterResult()
}

func userLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserLoginArgs)
	realResult := result.(*user.UserServiceUserLoginResult)
	success, err := handler.(user.UserService).UserLogin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserLoginArgs() interface{} {
	return user.NewUserServiceUserLoginArgs()
}

func newUserServiceUserLoginResult() interface{} {
	return user.NewUserServiceUserLoginResult()
}

func userInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserInfoArgs)
	realResult := result.(*user.UserServiceUserInfoResult)
	success, err := handler.(user.UserService).UserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserInfoArgs() interface{} {
	return user.NewUserServiceUserInfoArgs()
}

func newUserServiceUserInfoResult() interface{} {
	return user.NewUserServiceUserInfoResult()
}

func userInfoUploadHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserInfoUploadArgs)
	realResult := result.(*user.UserServiceUserInfoUploadResult)
	success, err := handler.(user.UserService).UserInfoUpload(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserInfoUploadArgs() interface{} {
	return user.NewUserServiceUserInfoUploadArgs()
}

func newUserServiceUserInfoUploadResult() interface{} {
	return user.NewUserServiceUserInfoUploadResult()
}

func userProfileInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserProfileInfoArgs)
	realResult := result.(*user.UserServiceUserProfileInfoResult)
	success, err := handler.(user.UserService).UserProfileInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserProfileInfoArgs() interface{} {
	return user.NewUserServiceUserProfileInfoArgs()
}

func newUserServiceUserProfileInfoResult() interface{} {
	return user.NewUserServiceUserProfileInfoResult()
}

func userProfileUploadHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserProfileUploadArgs)
	realResult := result.(*user.UserServiceUserProfileUploadResult)
	success, err := handler.(user.UserService).UserProfileUpload(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserProfileUploadArgs() interface{} {
	return user.NewUserServiceUserProfileUploadArgs()
}

func newUserServiceUserProfileUploadResult() interface{} {
	return user.NewUserServiceUserProfileUploadResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (r *user.UserRegisterResponse, err error) {
	var _args user.UserServiceUserRegisterArgs
	_args.Req = req
	var _result user.UserServiceUserRegisterResult
	if err = p.c.Call(ctx, "UserRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserLogin(ctx context.Context, req *user.UserLoginRequest) (r *user.UserLoginResponse, err error) {
	var _args user.UserServiceUserLoginArgs
	_args.Req = req
	var _result user.UserServiceUserLoginResult
	if err = p.c.Call(ctx, "UserLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserInfo(ctx context.Context, req *user.UserInfoRequest) (r *user.UserInfoResponse, err error) {
	var _args user.UserServiceUserInfoArgs
	_args.Req = req
	var _result user.UserServiceUserInfoResult
	if err = p.c.Call(ctx, "UserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserInfoUpload(ctx context.Context, req *user.UserInfoUploadRequest) (r *user.UserInfoUploadResponse, err error) {
	var _args user.UserServiceUserInfoUploadArgs
	_args.Req = req
	var _result user.UserServiceUserInfoUploadResult
	if err = p.c.Call(ctx, "UserInfoUpload", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserProfileInfo(ctx context.Context, req *user.UserProfileInfoRequest) (r *user.UserProfileInfoResponse, err error) {
	var _args user.UserServiceUserProfileInfoArgs
	_args.Req = req
	var _result user.UserServiceUserProfileInfoResult
	if err = p.c.Call(ctx, "UserProfileInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserProfileUpload(ctx context.Context, req *user.UserProfileUploadRequest) (r *user.UserProfileUploadResponse, err error) {
	var _args user.UserServiceUserProfileUploadArgs
	_args.Req = req
	var _result user.UserServiceUserProfileUploadResult
	if err = p.c.Call(ctx, "UserProfileUpload", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
