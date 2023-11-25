package rpc

import (
	"context"
	"github.com/Yra-A/Fusion_Go/kitex_gen/favorite"
	"github.com/Yra-A/Fusion_Go/kitex_gen/favorite/favoriteservice"
	"github.com/Yra-A/Fusion_Go/pkg/constants"
	"github.com/Yra-A/Fusion_Go/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var favoriteClient favoriteservice.Client

func initFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress}) // 服务发现
	if err != nil {
		panic(err)
	}
	c, err := favoriteservice.NewClient(
		constants.FavoriteServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

// ContestFavoriteAction 收藏比赛【rpc 客户端】
func ContestFavoriteAction(ctx context.Context, req *favorite.ContestFavoriteActionRequest) (*favorite.ContestFavoriteActionResponse, error) {
	resp, err := favoriteClient.ContestFavoriteAction(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ContestFavoriteList 获取收藏比赛列表【rpc 客户端】
func ContestFavoriteList(ctx context.Context, req *favorite.ContestFavoriteListRequest) (*favorite.ContestFavoriteListResponse, error) {
	resp, err := favoriteClient.ContestFavoriteList(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
