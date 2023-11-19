package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/contest/dal/db"
	"github.com/Yra-A/Fusion_Go/kitex_gen/contest"
	"sort"
)

type GetContestsByFavoritesService struct {
	ctx context.Context
}

func NewGetContestsByFavoritesService(ctx context.Context) *GetContestsByFavoritesService {
	return &GetContestsByFavoritesService{ctx: ctx}
}

func (s *GetContestsByFavoritesService) GetContestsByFavorites(contestIds []int32) ([]*contest.ContestBriefInfo, error) {
	dbContest, err := db.FetchContestListByContestIds(contestIds)
	if err != nil {
		return nil, err
	}
	// 创建排序映射
	sortOrder := make(map[int32]int)
	for i, id := range contestIds {
		sortOrder[id] = i
	}

	// 实现自定义排序, 按照contestIds的顺序排序
	sort.SliceStable(dbContest, func(i, j int) bool {
		return sortOrder[dbContest[i].ContestID] < sortOrder[dbContest[j].ContestID]
	})
	contestBriefInfos := make([]*contest.ContestBriefInfo, len(dbContest))
	for i, v := range dbContest {
		contestBriefInfos[i] = &contest.ContestBriefInfo{
			ContestBriefInfo: &contest.ContestBrief{
				ContestId:   v.ContestID,
				Title:       v.Title,
				Description: v.Description,
				CreatedTime: v.CreatedTime.Unix(),
				Field:       v.Field,
				Format:      v.Format,
			},
		}
	}
	return contestBriefInfos, nil
}
