package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/favorite/dal/db"
)

type QueryFavoriteStatusService struct {
	ctx context.Context
}

func NewQueryFavoriteStatusService(ctx context.Context) *QueryFavoriteStatusService {
	return &QueryFavoriteStatusService{ctx: ctx}
}

func (s *QueryFavoriteStatusService) QueryFavoriteStatusByUserId(userId int32, contestId int32) (isFavorite bool, err error) {
	isFavorite, err = db.QueryFavoriteStatusByUserId(userId, contestId)
	if err != nil {
		return false, err
	}
	return isFavorite, nil
}
