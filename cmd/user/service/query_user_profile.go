package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/user/dal/db"
	"github.com/Yra-A/Fusion_Go/kitex_gen/user"
	"sync"
)

type QueryUserProfileService struct {
	ctx context.Context
}

func NewQueryUserProfileService(ctx context.Context) *QueryUserProfileService {
	return &QueryUserProfileService{ctx: ctx}
}

func (s *QueryUserProfileService) QueryUserProfile(user_id int32) (*user.UserProfileInfo, error) {
	u := &user.UserProfileInfo{}
	tasks := []TaskFunc{
		func() error { return s.FetchUserProfileInfo(user_id, u) },
		//todo add func()...
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
	u.UserId = user_id
	return u, nil
}

func (s *QueryUserProfileService) FetchUserProfileInfo(user_id int32, u *user.UserProfileInfo) error {
	dbUserProfileInfo, err := db.QueryUserProfileByUserId(user_id)
	if err != nil {
		return err
	}
	u.UserId = dbUserProfileInfo.UserId
	u.MobilePhone = dbUserProfileInfo.MobilePhone
	u.Introduction = dbUserProfileInfo.Introduction
	u.QqNumber = dbUserProfileInfo.QqNumber
	u.WechatNumber = dbUserProfileInfo.WechatNumber
	u.Honors = dbUserProfileInfo.Honors
	u.Images = dbUserProfileInfo.Images
	u.IsShow = dbUserProfileInfo.IsShow
	return nil
}
