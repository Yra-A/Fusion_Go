package main

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/team/service"
	team "github.com/Yra-A/Fusion_Go/kitex_gen/team"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

// TeamServiceImpl implements the last service interface defined in the IDL.
type TeamServiceImpl struct{}

// TeamCreate implements the TeamServiceImpl interface.
func (s *TeamServiceImpl) TeamCreate(ctx context.Context, req *team.TeamCreateRequest) (resp *team.TeamCreateResponse, err error) {
	klog.CtxDebugf(ctx, "TeamCreate called")
	resp = new(team.TeamCreateResponse)
	team_id, err := service.NewCreateTeamService(ctx).CreateTeam(req.UserId, req.TeamId, req.ContestId, req.Title, req.Goal, req.Description)
	if err != nil {
		resp.StatusCode = errno.Fail.ErrCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, err
	}
	resp.TeamId = team_id
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	return resp, nil
}

// TeamList implements the TeamServiceImpl interface.
func (s *TeamServiceImpl) TeamList(ctx context.Context, req *team.TeamListRequest) (resp *team.TeamListResponse, err error) {
	klog.CtxDebugf(ctx, "TeamList called")
	resp = new(team.TeamListResponse)
	teamList, total, err := service.NewTeamListService(ctx).TeamList(req.ContestId, req.Limit, req.Offset)
	if err != nil {
		resp.StatusCode = errno.Fail.ErrCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, err
	}
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	resp.TeamList = teamList
	resp.Total = total
	return resp, nil
}

// TeamInfo implements the TeamServiceImpl interface.
func (s *TeamServiceImpl) TeamInfo(ctx context.Context, req *team.TeamInfoRequest) (resp *team.TeamInfoResponse, err error) {
	klog.CtxDebugf(ctx, "TeamInfo called")
	resp = new(team.TeamInfoResponse)
	teamInfo, err := service.NewTeamInfoService(ctx).TeamInfo(req.TeamId)
	if err != nil {
		resp.StatusCode = errno.Fail.ErrCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, err
	}
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	resp.TeamInfo = teamInfo
	return resp, nil
}

// TeamApplicationSubmit implements the TeamServiceImpl interface.
func (s *TeamServiceImpl) TeamApplicationSubmit(ctx context.Context, req *team.TeamApplicationSubmitRequest) (resp *team.TeamApplicationSubmitResponse, err error) {
	klog.CtxDebugf(ctx, "TeamApplicationSubmit called")
	resp = new(team.TeamApplicationSubmitResponse)
	err = service.NewTeamApplicationSubmitService(ctx).TeamApplicationSubmit(req.TeamId, req.Reason, req.CreatedTime, req.ApplicationType, req.MemberInfo.UserId)
	if err != nil {
		resp.StatusCode = errno.Fail.ErrCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, err
	}
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	return resp, nil
}

// TeamManageList implements the TeamServiceImpl interface.
func (s *TeamServiceImpl) TeamManageList(ctx context.Context, req *team.TeamManageListRequest) (resp *team.TeamManageListResponse, err error) {
	klog.CtxDebugf(ctx, "TeamManageList called")
	resp = new(team.TeamManageListResponse)
	teamManageList, err := service.NewTeamManageListService(ctx).TeamManageList(req.UserId, req.TeamId)
	if err != nil {
		resp.StatusCode = errno.Fail.ErrCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, err
	}
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	resp.ApplicationList = teamManageList
	return
}

// TeamManageAction implements the TeamServiceImpl interface.
func (s *TeamServiceImpl) TeamManageAction(ctx context.Context, req *team.TeamManageActionRequest) (resp *team.TeamManageActionResponse, err error) {
	klog.CtxDebugf(ctx, "TeamManageAction called")
	resp = new(team.TeamManageActionResponse)
	err = service.NewTeamManageActionService(ctx).TeamManageAction(req.UserId, req.ApplicationId, req.ActionType)
	if err != nil {
		resp.StatusCode = errno.Fail.ErrCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, err
	}
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	return
}
