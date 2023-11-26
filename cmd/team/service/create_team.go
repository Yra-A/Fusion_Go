package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/team/dal/db"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
)

type CreateTeamService struct {
	ctx context.Context
}

func NewCreateTeamService(ctx context.Context) *CreateTeamService {
	return &CreateTeamService{ctx: ctx}
}
func (s *CreateTeamService) CreateTeam(user_id int32, team_id int32, contest_id int32, title string, goal string, description string) (int32, error) {
	// team_id == 0 代表创建团队
	if team_id == 0 {
		return db.CreateTeam(user_id, contest_id, title, goal, description)
	} else if team_id > 0 {
		return team_id, db.ModifyTeam(team_id, title, goal, description)
	}
	return 0, errno.ParamErr
}
