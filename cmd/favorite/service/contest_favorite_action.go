package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/favorite/dal/db"
)

type ContestFavoriteActionService struct {
	ctx context.Context
}

func NewContestFavoriteActionService(ctx context.Context) *ContestFavoriteActionService {
	return &ContestFavoriteActionService{ctx: ctx}
}

func (s *ContestFavoriteActionService) ContestFavoriteAction(user_id int32, contest_id int32, action_type int32) error {
	err := db.ContestFavoriteAction(user_id, contest_id, action_type)
	if err != nil {
		return err
	}
	return nil
}
