package main

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/contest/service"
	contest "github.com/Yra-A/Fusion_Go/kitex_gen/contest"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

// ContestServiceImpl implements the last service interface defined in the IDL.
type ContestServiceImpl struct{}

// ContestCreate implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) ContestCreate(ctx context.Context, req *contest.ContestCreateRequest) (resp *contest.ContestCreateResponse, err error) {
	klog.CtxDebugf(ctx, "ContestCreate called: %v", req.GetContest().ContestId)
	resp = new(contest.ContestCreateResponse)
	contest_id, err := service.NewCreateContestService(ctx).CreateContest(req.Contest)
	if err == errno.ContestNotExistErr {
		resp.StatusCode = errno.ContestNotExistErr.ErrCode
		resp.StatusMsg = errno.ContestNotExistErr.ErrMsg
		return resp, nil
	}
	if err != nil {
		resp.StatusCode = errno.Fail.ErrCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, err
	}
	resp.ContestId = contest_id
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	return resp, nil
}

// ContestList implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) ContestList(ctx context.Context, req *contest.ContestListRequest) (resp *contest.ContestListResponse, err error) {
	klog.CtxDebugf(ctx, "ContestList called")
	resp = new(contest.ContestListResponse)
	c, total, err := service.NewQueryContestListService(ctx).QueryContestList(req.Keyword, req.Fields, req.Formats, req.Limit, req.Offset)
	if err != nil {
		resp.StatusCode = errno.Fail.ErrCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, nil
	}
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	resp.Total = total
	resp.ContestList = c
	return resp, nil
}

// ContestInfo implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) ContestInfo(ctx context.Context, req *contest.ContestInfoRequest) (resp *contest.ContestInfoResponse, err error) {
	klog.CtxDebugf(ctx, "ContestInfo called: %v", req.GetContestId())
	resp = new(contest.ContestInfoResponse)
	c, err := service.NewQueryContestService(ctx).QueryContest(req.UserId, req.ContestId)
	if err != nil {
		resp.StatusCode = errno.Fail.ErrCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, nil
	}
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	resp.Contest = c
	return resp, nil
}

// GetContestsByFavorites implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) GetContestsByFavorites(ctx context.Context, req *contest.GetContestsByFavoritesRequest) (resp *contest.GetContestsByFavoritesResponse, err error) {
	resp = new(contest.GetContestsByFavoritesResponse)
	c, err := service.NewGetContestsByFavoritesService(ctx).GetContestsByFavorites(req.ContestIds)
	if err != nil {
		return nil, err
	}
	resp.ContestList = c
	return resp, nil
}
