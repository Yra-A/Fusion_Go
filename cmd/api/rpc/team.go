package rpc

import (
	"context"
	"github.com/Yra-A/Fusion_Go/kitex_gen/team"
	"github.com/Yra-A/Fusion_Go/kitex_gen/team/teamservice"
	"github.com/Yra-A/Fusion_Go/pkg/constants"
	"github.com/Yra-A/Fusion_Go/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var teamClient teamservice.Client

func initTeamRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress}) // 服务发现
	if err != nil {
		panic(err)
	}
	c, err := teamservice.NewClient(
		constants.TeamServiceName,
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
	teamClient = c
}

// TeamCreate 创建团队【rpc 客户端】
func TeamCreate(ctx context.Context, req *team.TeamCreateRequest) (*team.TeamCreateResponse, error) {
	resp, err := teamClient.TeamCreate(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TeamList 获取团队列表【rpc 客户端】
func TeamList(ctx context.Context, req *team.TeamListRequest) (*team.TeamListResponse, error) {
	resp, err := teamClient.TeamList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TeamInfo 获取团队信息【rpc 客户端】
func TeamInfo(ctx context.Context, req *team.TeamInfoRequest) (*team.TeamInfoResponse, error) {
	resp, err := teamClient.TeamInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TeamApplicationSubmit 提交团队申请【rpc 客户端】
func TeamApplicationSubmit(ctx context.Context, req *team.TeamApplicationSubmitRequest) (*team.TeamApplicationSubmitResponse, error) {
	resp, err := teamClient.TeamApplicationSubmit(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TeamManageList 获取团队收到的申请列表【rpc 客户端】
func TeamManageList(ctx context.Context, req *team.TeamManageListRequest) (*team.TeamManageListResponse, error) {
	resp, err := teamClient.TeamManageList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TeamManageAction 团队申请处理【rpc 客户端】
func TeamManageAction(ctx context.Context, req *team.TeamManageActionRequest) (*team.TeamManageActionResponse, error) {
	resp, err := teamClient.TeamManageAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
