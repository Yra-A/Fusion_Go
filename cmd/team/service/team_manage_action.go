package service

import (
    "context"
    "github.com/Yra-A/Fusion_Go/cmd/team/dal/db"
)

type TeamManageActionService struct {
    ctx context.Context
}

func NewTeamManageActionService(ctx context.Context) *TeamManageActionService {
    return &TeamManageActionService{ctx: ctx}
}

func (s *TeamManageActionService) TeamManageAction(user_id int32, application_id int32, action_type int32) error {
    err := db.TeamManageAction(user_id, application_id, action_type)
    if err != nil {
        return err
    }
    return nil
}
