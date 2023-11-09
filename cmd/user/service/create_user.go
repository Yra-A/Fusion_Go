package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/user/dal/db"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
	"github.com/Yra-A/Fusion_Go/pkg/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}
func (s *CreateUserService) CreateUser(username, password string) error {
	var err error
	if _, err = db.QueryUserByName(username); err == nil {
		return errno.UserAlreadyExistErr
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hash, err := utils.PasswordHash(password)
	if err != nil {
		return err
	}
	return db.CreateUser(username, hash)
}
