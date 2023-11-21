package main

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/favorite/service"
	favorite "github.com/Yra-A/Fusion_Go/kitex_gen/favorite"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// ContestFavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) ContestFavoriteAction(ctx context.Context, req *favorite.ContestFavoriteActionRequest) (resp *favorite.ContestFavoriteActionResponse, err error) {
	klog.CtxDebugf(ctx, "ContestFavoriteAction called: %v", req.GetContestId())
	resp = new(favorite.ContestFavoriteActionResponse)
	err = service.NewContestFavoriteActionService(ctx).ContestFavoriteAction(req.UserId, req.ContestId, req.ActionType)
	if err != nil {
		resp.StatusCode = errno.Fail.ErrCode
		resp.StatusMsg = err.Error()
		return resp, nil
	}
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	return resp, nil
}

// ContestFavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) ContestFavoriteList(ctx context.Context, req *favorite.ContestFavoriteListRequest) (resp *favorite.ContestFavoriteListResponse, err error) {
	klog.CtxDebugf(ctx, "ContestFavoriteList called: %v", req.GetUserId())
	resp = new(favorite.ContestFavoriteListResponse)
	c, err := service.NewQueryContestFavoriteListService(ctx).QueryContestFavoriteList(req.UserId, req.Limit, req.Offset)
	if err != nil {
		resp.StatusCode = errno.Fail.ErrCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, nil
	}
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	resp.ContestList = c
	resp.Total = int32(len(c))
	return resp, nil
}

// QueryFavoriteStatusByUserId implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) QueryFavoriteStatusByUserId(ctx context.Context, req *favorite.QueryFavoriteStatusByUserIdRequest) (resp *favorite.QueryFavoriteStatusByUserIdResponse, err error) {
	klog.CtxDebugf(ctx, "QueryFavoriteStatusByUserId called: %v", req.GetUserId())
	resp = new(favorite.QueryFavoriteStatusByUserIdResponse)
	isFavorite, err := service.NewQueryFavoriteStatusService(ctx).QueryFavoriteStatusByUserId(req.UserId, req.ContestId)
	if err != nil {
		return nil, err
	}
	resp.IsFavorite = isFavorite
	return resp, nil
}
