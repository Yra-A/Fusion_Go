package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/user/dal/db"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}
func (s *CheckUserService) CheckUser(username, password string) (int32, error) {
	user, err := db.QueryUserByName(username)
	if err != nil {
		return 0, err
	}
	if user == nil {
		return 0, errno.UserNotExistErr
	}
	if user.Password != password {
		return 0, errno.InvalidCredentialsErr
	}

	return user.UserID, nil
}
