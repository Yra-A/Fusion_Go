package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/contest/dal/db"
	"github.com/Yra-A/Fusion_Go/kitex_gen/contest"
)

type QueryContestListService struct {
	ctx context.Context
}

func NewQueryContestListService(ctx context.Context) *QueryContestListService {
	return &QueryContestListService{ctx: ctx}
}
func (s *QueryContestListService) QueryContestList(keyword string, fields []string, formats []string, limit int32, offset int32) ([]*contest.ContestBriefInfo, int32, error) {
	dbContests, total, err := db.FetchContestList(keyword, fields, formats, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	contestBriefInfos := make([]*contest.ContestBriefInfo, len(dbContests))
	for i, v := range dbContests {
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
	return contestBriefInfos, total, nil
}
