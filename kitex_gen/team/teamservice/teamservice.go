// Code generated by Kitex v0.6.2. DO NOT EDIT.

package teamservice

import (
	"context"
	team "github.com/Yra-A/Fusion_Go/kitex_gen/team"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return teamServiceServiceInfo
}

var teamServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "TeamService"
	handlerType := (*team.TeamService)(nil)
	methods := map[string]kitex.MethodInfo{
		"TeamCreate":            kitex.NewMethodInfo(teamCreateHandler, newTeamServiceTeamCreateArgs, newTeamServiceTeamCreateResult, false),
		"TeamList":              kitex.NewMethodInfo(teamListHandler, newTeamServiceTeamListArgs, newTeamServiceTeamListResult, false),
		"TeamInfo":              kitex.NewMethodInfo(teamInfoHandler, newTeamServiceTeamInfoArgs, newTeamServiceTeamInfoResult, false),
		"TeamApplicationSubmit": kitex.NewMethodInfo(teamApplicationSubmitHandler, newTeamServiceTeamApplicationSubmitArgs, newTeamServiceTeamApplicationSubmitResult, false),
		"TeamManageList":        kitex.NewMethodInfo(teamManageListHandler, newTeamServiceTeamManageListArgs, newTeamServiceTeamManageListResult, false),
		"TeamManageAction":      kitex.NewMethodInfo(teamManageActionHandler, newTeamServiceTeamManageActionArgs, newTeamServiceTeamManageActionResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "team",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func teamCreateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*team.TeamServiceTeamCreateArgs)
	realResult := result.(*team.TeamServiceTeamCreateResult)
	success, err := handler.(team.TeamService).TeamCreate(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTeamServiceTeamCreateArgs() interface{} {
	return team.NewTeamServiceTeamCreateArgs()
}

func newTeamServiceTeamCreateResult() interface{} {
	return team.NewTeamServiceTeamCreateResult()
}

func teamListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*team.TeamServiceTeamListArgs)
	realResult := result.(*team.TeamServiceTeamListResult)
	success, err := handler.(team.TeamService).TeamList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTeamServiceTeamListArgs() interface{} {
	return team.NewTeamServiceTeamListArgs()
}

func newTeamServiceTeamListResult() interface{} {
	return team.NewTeamServiceTeamListResult()
}

func teamInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*team.TeamServiceTeamInfoArgs)
	realResult := result.(*team.TeamServiceTeamInfoResult)
	success, err := handler.(team.TeamService).TeamInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTeamServiceTeamInfoArgs() interface{} {
	return team.NewTeamServiceTeamInfoArgs()
}

func newTeamServiceTeamInfoResult() interface{} {
	return team.NewTeamServiceTeamInfoResult()
}

func teamApplicationSubmitHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*team.TeamServiceTeamApplicationSubmitArgs)
	realResult := result.(*team.TeamServiceTeamApplicationSubmitResult)
	success, err := handler.(team.TeamService).TeamApplicationSubmit(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTeamServiceTeamApplicationSubmitArgs() interface{} {
	return team.NewTeamServiceTeamApplicationSubmitArgs()
}

func newTeamServiceTeamApplicationSubmitResult() interface{} {
	return team.NewTeamServiceTeamApplicationSubmitResult()
}

func teamManageListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*team.TeamServiceTeamManageListArgs)
	realResult := result.(*team.TeamServiceTeamManageListResult)
	success, err := handler.(team.TeamService).TeamManageList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTeamServiceTeamManageListArgs() interface{} {
	return team.NewTeamServiceTeamManageListArgs()
}

func newTeamServiceTeamManageListResult() interface{} {
	return team.NewTeamServiceTeamManageListResult()
}

func teamManageActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*team.TeamServiceTeamManageActionArgs)
	realResult := result.(*team.TeamServiceTeamManageActionResult)
	success, err := handler.(team.TeamService).TeamManageAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTeamServiceTeamManageActionArgs() interface{} {
	return team.NewTeamServiceTeamManageActionArgs()
}

func newTeamServiceTeamManageActionResult() interface{} {
	return team.NewTeamServiceTeamManageActionResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) TeamCreate(ctx context.Context, req *team.TeamCreateRequest) (r *team.TeamCreateResponse, err error) {
	var _args team.TeamServiceTeamCreateArgs
	_args.Req = req
	var _result team.TeamServiceTeamCreateResult
	if err = p.c.Call(ctx, "TeamCreate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TeamList(ctx context.Context, req *team.TeamListRequest) (r *team.TeamListResponse, err error) {
	var _args team.TeamServiceTeamListArgs
	_args.Req = req
	var _result team.TeamServiceTeamListResult
	if err = p.c.Call(ctx, "TeamList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TeamInfo(ctx context.Context, req *team.TeamInfoRequest) (r *team.TeamInfoResponse, err error) {
	var _args team.TeamServiceTeamInfoArgs
	_args.Req = req
	var _result team.TeamServiceTeamInfoResult
	if err = p.c.Call(ctx, "TeamInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TeamApplicationSubmit(ctx context.Context, req *team.TeamApplicationSubmitRequest) (r *team.TeamApplicationSubmitResponse, err error) {
	var _args team.TeamServiceTeamApplicationSubmitArgs
	_args.Req = req
	var _result team.TeamServiceTeamApplicationSubmitResult
	if err = p.c.Call(ctx, "TeamApplicationSubmit", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TeamManageList(ctx context.Context, req *team.TeamManageListRequest) (r *team.TeamManageListResponse, err error) {
	var _args team.TeamServiceTeamManageListArgs
	_args.Req = req
	var _result team.TeamServiceTeamManageListResult
	if err = p.c.Call(ctx, "TeamManageList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TeamManageAction(ctx context.Context, req *team.TeamManageActionRequest) (r *team.TeamManageActionResponse, err error) {
	var _args team.TeamServiceTeamManageActionArgs
	_args.Req = req
	var _result team.TeamServiceTeamManageActionResult
	if err = p.c.Call(ctx, "TeamManageAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
