package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/team/dal/db"
	"github.com/Yra-A/Fusion_Go/cmd/team/rpc"
	"github.com/Yra-A/Fusion_Go/kitex_gen/team"
	"github.com/Yra-A/Fusion_Go/kitex_gen/user"
)

type TeamListService struct {
	ctx context.Context
}

func NewTeamListService(ctx context.Context) *TeamListService {
	return &TeamListService{ctx: ctx}
}

func (s *TeamListService) TeamList(contest_id int32, limit int32, offset int32) ([]*team.TeamBriefInfo, int32, error) {
	teamList, err := db.QueryTeamList(contest_id)
	if err != nil {
		return nil, 0, err
	}
	total := len(teamList)
	if offset < int32(len(teamList)) {
		if offset+limit >= int32(len(teamList)) {
			teamList = teamList[offset:]
		} else {
			teamList = teamList[offset : offset+limit]
		}
	} else {
		teamList = nil
	}

	for _, t := range teamList {
		kresp, err := rpc.UserProfileInfo(context.Background(), &user.UserProfileInfoRequest{UserId: t.LeaderInfo.UserId})
		if err != nil {
			return nil, 0, err
		}
		u := kresp.UserProfileInfo
		t.LeaderInfo = &team.MemberInfo{
			UserId:         u.UserInfo.UserId,
			Nickname:       u.UserInfo.Nickname,
			AvatarUrl:      u.UserInfo.AvatarUrl,
			College:        u.UserInfo.College,
			EnrollmentYear: u.UserInfo.EnrollmentYear,
			Gender:         u.UserInfo.Gender,
			Honors:         u.Honors,
		}
	}
	return teamList, int32(total), nil
}
