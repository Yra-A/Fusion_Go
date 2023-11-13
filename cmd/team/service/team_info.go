package service

import (
    "context"
    "github.com/Yra-A/Fusion_Go/cmd/team/dal/db"
    "github.com/Yra-A/Fusion_Go/cmd/team/rpc"
    "github.com/Yra-A/Fusion_Go/kitex_gen/team"
    "github.com/Yra-A/Fusion_Go/kitex_gen/user"
)

type TeamInfoService struct {
    ctx context.Context
}

func NewTeamInfoService(ctx context.Context) *TeamInfoService {
    return &TeamInfoService{ctx: ctx}
}

func convertUserProfileInfoToMemberInfo(u *user.UserProfileInfo) *team.MemberInfo {
    return &team.MemberInfo{
        UserId:         u.UserInfo.UserId,
        Nickname:       u.UserInfo.Nickname,
        AvatarUrl:      u.UserInfo.AvatarUrl,
        College:        u.UserInfo.College,
        EnrollmentYear: u.UserInfo.EnrollmentYear,
        Gender:         u.UserInfo.Gender,
        Honors:         u.Honors,
    }
}

func (s *TeamInfoService) TeamInfo(team_id int32) (*team.TeamInfo, error) {
    teamInfo, err := db.QueryTeamInfo(team_id)
    if err != nil {
        return nil, err
    }

    kresp, err := rpc.UserProfileInfo(context.Background(), &user.UserProfileInfoRequest{UserId: teamInfo.TeamBriefInfo.LeaderInfo.UserId})
    if err != nil {
        return nil, err
    }
    teamInfo.TeamBriefInfo.LeaderInfo = convertUserProfileInfoToMemberInfo(kresp.UserProfileInfo)

    for i, t := range teamInfo.Members {
        kresp, err := rpc.UserProfileInfo(context.Background(), &user.UserProfileInfoRequest{UserId: t.UserId})
        if err != nil {
            return nil, err
        }
        u := convertUserProfileInfoToMemberInfo(kresp.UserProfileInfo)
        teamInfo.Members[i] = u
    }
    return teamInfo, nil
}
