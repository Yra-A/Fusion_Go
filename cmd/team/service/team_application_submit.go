package service

import (
    "context"
    "github.com/Yra-A/Fusion_Go/cmd/team/dal/db"
)

type TeamApplicationSubmitService struct {
    ctx context.Context
}

func NewTeamApplicationSubmitService(ctx context.Context) *TeamApplicationSubmitService {
    return &TeamApplicationSubmitService{ctx: ctx}
}

func (s *TeamApplicationSubmitService) TeamApplicationSubmit(team_id int32, reason string, created_time int64, application_type int32, user_id int32) error {
    err := db.CreateTeamApplication(user_id, team_id, reason, created_time, application_type)
    if err != nil {
        return err
    }
    return nil
}
