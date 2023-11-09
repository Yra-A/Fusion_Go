package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/user/dal/db"
	"github.com/Yra-A/Fusion_Go/kitex_gen/user"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
)

type TaskFunc func() error

type QueryUserService struct {
	ctx context.Context
}

func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{ctx: ctx}
}

func (s *QueryUserService) QueryUser(user_id int32) (*user.UserInfo, error) {
	u := &user.UserInfo{}

	tasks := []TaskFunc{
		func() error { return s.FetchUserInfo(user_id, u) },
		//todo add func()
	}

	errChan := make(chan error, len(tasks))
	defer close(errChan)
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t TaskFunc) {
			defer wg.Done()
			if err := t(); err != nil {
				errChan <- err
			}
		}(task)
	}
	wg.Wait()
	select {
	case err := <-errChan:
		return nil, err
	default:
	}
	return u, nil
}
func (s *QueryUserService) FetchUserInfo(user_id int32, u *user.UserInfo) error {
	dbUserInfo, err := db.QueryUserByUserId(user_id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errno.UserNotExistErr
	}
	if err != nil {
		return err
	}

	u.UserId = dbUserInfo.UserID
	u.Gender = dbUserInfo.Gender
	u.EnrollmentYear = dbUserInfo.EnrollmentYear
	u.MobilePhone = dbUserInfo.MobilePhone
	u.College = dbUserInfo.College
	u.Nickname = dbUserInfo.Nickname
	u.Realname = dbUserInfo.Realname
	u.HasProfile = dbUserInfo.HasProfile
	u.AvatarUrl = dbUserInfo.AvatarURL

	return nil
}
