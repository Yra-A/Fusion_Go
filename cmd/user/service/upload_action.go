package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/user/dal/db"
	"github.com/Yra-A/Fusion_Go/kitex_gen/user"
)

type UploadUserService struct {
	ctx context.Context
}

func NewUploadUserService(ctx context.Context) *UploadUserService {
	return &UploadUserService{ctx: ctx}
}

func (s *UploadUserService) UploadUserInfo(u *user.UserInfo) error {
	dbu := &db.UserInfo{
		UserId:               u.UserId,
		Gender:               u.Gender,
		Nickname:             u.Nickname,
		Realname:             u.Realname,
		ContestFavoriteCount: u.ContestFavoriteCount,
		AvatarUrl:            u.AvatarUrl,
		EnrollmentYear:       u.EnrollmentYear,
		College:              u.College,
	}
	if err := db.AddUserInfo(dbu); err != nil {
		return err
	}
	return nil
}

func (s *UploadUserService) UploadUserProfileInfo(u *user.UserProfileInfo) error {
	dbu := &db.UserProfileInfo{
		UserId:       u.UserId,
		MobilePhone:  u.MobilePhone,
		Introduction: u.Introduction,
		QqNumber:     u.QqNumber,
		WechatNumber: u.WechatNumber,
		Honors:       u.Honors,
		Images:       u.Images,
		IsShow:       u.IsShow,
	}
	if err := db.AddUserProfileInfo(dbu); err != nil {
		return err
	}
	return nil
}
