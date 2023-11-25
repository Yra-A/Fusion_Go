package rpc

import (
	"context"
	"github.com/Yra-A/Fusion_Go/kitex_gen/contest"
	"github.com/Yra-A/Fusion_Go/kitex_gen/contest/contestservice"
	"github.com/Yra-A/Fusion_Go/pkg/constants"
	"github.com/Yra-A/Fusion_Go/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var contestClient contestservice.Client

func initContestRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress}) // 服务发现
	if err != nil {
		panic(err)
	}
	c, err := contestservice.NewClient(
		constants.ContestServiceName,
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
	contestClient = c
}

// ContestCreate 创建赛事【rpc 客户端】
func ContestCreate(ctx context.Context, req *contest.ContestCreateRequest) (*contest.ContestCreateResponse, error) {
	resp, err := contestClient.ContestCreate(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ContestList 比赛列表【rpc 客户端】
func ContestList(ctx context.Context, req *contest.ContestListRequest) (*contest.ContestListResponse, error) {
	resp, err := contestClient.ContestList(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ContestInfo 比赛详情【rpc 客户端】
func ContestInfo(ctx context.Context, req *contest.ContestInfoRequest) (*contest.ContestInfoResponse, error) {
	resp, err := contestClient.ContestInfo(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
