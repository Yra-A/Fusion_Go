package rpc

import (
	"context"
	"github.com/Yra-A/Fusion_Go/kitex_gen/user"
	"github.com/Yra-A/Fusion_Go/kitex_gen/user/userservice"
	"github.com/Yra-A/Fusion_Go/pkg/constants"
	"github.com/Yra-A/Fusion_Go/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress}) // 服务发现
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(tracing.NewClientSuite()),        // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// UserRegister 用户注册【rpc 客户端】
func UserRegister(ctx context.Context, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	resp, err := userClient.UserRegister(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// UserLogin 用户登录【rpc 客户端】
func UserLogin(ctx context.Context, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	resp, err := userClient.UserLogin(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// UserInfo 用户信息【rpc 客户端】
func UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	resp, err := userClient.UserInfo(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// UserInfoUpload 用户信息上传【rpc 客户端】
func UserInfoUpload(ctx context.Context, req *user.UserInfoUploadRequest) (*user.UserInfoUploadResponse, error) {
	resp, err := userClient.UserInfoUpload(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// UserProfileInfo 用户资料信息【rpc 客户端】
func UserProfileInfo(ctx context.Context, req *user.UserProfileInfoRequest) (*user.UserProfileInfoResponse, error) {
	resp, err := userClient.UserProfileInfo(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// UserProfileUpload 用户资料上传【rpc 客户端】
func UserProfileUpload(ctx context.Context, req *user.UserProfileUploadRequest) (*user.UserProfileUploadResponse, error) {
	resp, err := userClient.UserProfileUpload(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
