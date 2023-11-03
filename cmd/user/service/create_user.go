package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/user/dal/db"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
	"github.com/Yra-A/Fusion_Go/pkg/utils"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}
func (s *CreateUserService) CreateUser(username, password string) error {
	u, err := db.QueryUserByName(username)
	if err != nil {
		return err
	}

	if u != nil {
		return errno.UserAlreadyExistErr
	}
	hash, err := utils.PasswordHash(password)
	if err != nil {
		return err
	}
	err = db.CreateUser(username, hash)
	if err != nil {
		return err
	}
	return nil
}
