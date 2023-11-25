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
		func() error { return s.FetchUserHonors(user_id, u) },
		func() error { return s.FetchUserInfo(user_id, u) },
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

func (s *QueryUserProfileService) FetchUserProfileInfo(user_id int32, u *user.UserProfileInfo) error {
	dbUserProfileInfo, err := db.QueryUserProfileByUserId(db.DB, user_id)
	if err != nil {
		return err
	}
	u.Introduction = dbUserProfileInfo.Introduction
	u.QqNumber = dbUserProfileInfo.QQNumber
	u.WechatNumber = dbUserProfileInfo.WeChatNumber
	return nil
}

func (s *QueryUserProfileService) FetchUserHonors(user_id int32, u *user.UserProfileInfo) error {
	dbHonors, err := db.QueryHonorsByUserId(user_id)
	if err != nil {
		return err
	}
	u.Honors = dbHonors
	return nil
}

func (s *QueryUserProfileService) FetchUserInfo(user_id int32, u *user.UserProfileInfo) error {
	dbUserInfo, err := db.QueryUserByUserId(user_id)
	if err != nil {
		return err
	}
	u.UserInfo = &user.UserInfo{
		UserId:         dbUserInfo.UserID,
		Gender:         dbUserInfo.Gender,
		EnrollmentYear: dbUserInfo.EnrollmentYear,
		MobilePhone:    dbUserInfo.MobilePhone,
		College:        dbUserInfo.College,
		Nickname:       dbUserInfo.Nickname,
		Realname:       dbUserInfo.Realname,
		HasProfile:     dbUserInfo.HasProfile,
		AvatarUrl:      dbUserInfo.AvatarURL,
	}
	return nil
}
