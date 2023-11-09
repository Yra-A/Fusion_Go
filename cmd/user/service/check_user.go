package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/user/dal/db"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
	"github.com/Yra-A/Fusion_Go/pkg/utils"
	"gorm.io/gorm"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}
func (s *CheckUserService) CheckUser(username, password string) (int32, error) {
	var err error
	user, err := db.QueryUserByName(username)
	if err == gorm.ErrRecordNotFound {
		return 0, errno.UserNotExistErr
	}
	if err != nil {
		return 0, err
	}
	if ok := utils.PasswordVerify(password, user.Password); !ok {
		return 0, errno.InvalidCredentialsErr
	}
	return user.UserID, nil
}
