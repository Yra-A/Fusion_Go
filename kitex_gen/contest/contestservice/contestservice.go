// Code generated by Kitex v0.7.3. DO NOT EDIT.

package contestservice

import (
	"context"
	contest "github.com/Yra-A/Fusion_Go/kitex_gen/contest"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return contestServiceServiceInfo
}

var contestServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ContestService"
	handlerType := (*contest.ContestService)(nil)
	methods := map[string]kitex.MethodInfo{
		"ContestList":   kitex.NewMethodInfo(contestListHandler, newContestServiceContestListArgs, newContestServiceContestListResult, false),
		"ContestInfo":   kitex.NewMethodInfo(contestInfoHandler, newContestServiceContestInfoArgs, newContestServiceContestInfoResult, false),
		"ContestCreate": kitex.NewMethodInfo(contestCreateHandler, newContestServiceContestCreateArgs, newContestServiceContestCreateResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "contest",
		"ServiceFilePath": `idl/contest.thrift`,
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

func contestListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*contest.ContestServiceContestListArgs)
	realResult := result.(*contest.ContestServiceContestListResult)
	success, err := handler.(contest.ContestService).ContestList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newContestServiceContestListArgs() interface{} {
	return contest.NewContestServiceContestListArgs()
}

func newContestServiceContestListResult() interface{} {
	return contest.NewContestServiceContestListResult()
}

func contestInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*contest.ContestServiceContestInfoArgs)
	realResult := result.(*contest.ContestServiceContestInfoResult)
	success, err := handler.(contest.ContestService).ContestInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newContestServiceContestInfoArgs() interface{} {
	return contest.NewContestServiceContestInfoArgs()
}

func newContestServiceContestInfoResult() interface{} {
	return contest.NewContestServiceContestInfoResult()
}

func contestCreateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*contest.ContestServiceContestCreateArgs)
	realResult := result.(*contest.ContestServiceContestCreateResult)
	success, err := handler.(contest.ContestService).ContestCreate(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newContestServiceContestCreateArgs() interface{} {
	return contest.NewContestServiceContestCreateArgs()
}

func newContestServiceContestCreateResult() interface{} {
	return contest.NewContestServiceContestCreateResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ContestList(ctx context.Context, req *contest.ContestListRequest) (r *contest.ContestListResponse, err error) {
	var _args contest.ContestServiceContestListArgs
	_args.Req = req
	var _result contest.ContestServiceContestListResult
	if err = p.c.Call(ctx, "ContestList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContestInfo(ctx context.Context, req *contest.ContestInfoRequest) (r *contest.ContestInfoResponse, err error) {
	var _args contest.ContestServiceContestInfoArgs
	_args.Req = req
	var _result contest.ContestServiceContestInfoResult
	if err = p.c.Call(ctx, "ContestInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContestCreate(ctx context.Context, req *contest.ContestCreateRequest) (r *contest.ContestCreateResponse, err error) {
	var _args contest.ContestServiceContestCreateArgs
	_args.Req = req
	var _result contest.ContestServiceContestCreateResult
	if err = p.c.Call(ctx, "ContestCreate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
