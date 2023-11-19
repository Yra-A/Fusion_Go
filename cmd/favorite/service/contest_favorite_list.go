package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/favorite/dal/db"
	"github.com/Yra-A/Fusion_Go/cmd/favorite/rpc"
	"github.com/Yra-A/Fusion_Go/kitex_gen/contest"
	"github.com/Yra-A/Fusion_Go/kitex_gen/favorite"
	"github.com/Yra-A/Fusion_Go/pkg/utils"
)

type QueryContestFavoriteListService struct {
	ctx context.Context
}

func NewQueryContestFavoriteListService(ctx context.Context) *QueryContestFavoriteListService {
	return &QueryContestFavoriteListService{ctx: ctx}
}

func (s *QueryContestFavoriteListService) QueryContestFavoriteList(user_id int32, limit int32, offset int32) ([]*favorite.ContestBriefInfo, error) {
	var err error
	contestIDs, err := db.FetchContestFavoriteList(user_id, limit, offset)
	if err != nil {
		return nil, err
	}
	kresp, err := rpc.GetContestsByFavorites(s.ctx, &contest.GetContestsByFavoritesRequest{ContestIds: contestIDs})
	if err != nil {
		return nil, err
	}
	contestBriefInfos := make([]*favorite.ContestBriefInfo, len(kresp.ContestList))
	for i, v := range kresp.ContestList {
		contestBriefInfos[i] = &favorite.ContestBriefInfo{
			ContestBriefInfo: utils.ConvertContestBriefToFavoriteBrief(v.ContestBriefInfo),
		}
	}
	return contestBriefInfos, nil
}
