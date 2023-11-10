package service

import (
    "context"
    "github.com/Yra-A/Fusion_Go/cmd/team/dal/db"
    "github.com/Yra-A/Fusion_Go/cmd/team/rpc"
    "github.com/Yra-A/Fusion_Go/kitex_gen/team"
    "github.com/Yra-A/Fusion_Go/kitex_gen/user"
)

type TeamManageListService struct {
    ctx context.Context
}

func NewTeamManageListService(ctx context.Context) *TeamManageListService {
    return &TeamManageListService{ctx: ctx}
}

func (s *TeamManageListService) TeamManageList(user_id int32, team_id int32) ([]*team.TeamApplication, error) {
    teamApplicationList, err := db.GetTeamApplicationList(user_id, team_id)
    if err != nil {
        return nil, err
    }
    for _, t := range teamApplicationList {
        kresp, err := rpc.UserProfileInfo(context.Background(), &user.UserProfileInfoRequest{UserId: t.MemberInfo.UserId})
        if err != nil {
            return nil, err
        }
        u := convertUserProfileInfoToMemberInfo(kresp.UserProfileInfo)
        t.MemberInfo = u
    }

    return teamApplicationList, nil
}
