package main

import (
	"context"
	user "github.com/Yra-A/Fusion_Go/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// UserInfoUpload implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfoUpload(ctx context.Context, req *user.UserInfoUploadRequest) (resp *user.UserInfoUploadResponse, err error) {
	// TODO: Your code here...
	return
}

// UserProfileInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserProfileInfo(ctx context.Context, req *user.UserProfileInfoRequest) (resp *user.UserProfileInfoResponse, err error) {
	resp = new(user.UserProfileInfoResponse)

	return
}

// UserProfileUpload implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserProfileUpload(ctx context.Context, req *user.UserProfileUploadRequest) (resp *user.UserProfileUploadResponse, err error) {
	// TODO: Your code here...
	return
}
