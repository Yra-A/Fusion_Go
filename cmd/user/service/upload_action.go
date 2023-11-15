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
	dbu := &db.UserProfileInfo{
		UserID:         u.UserId,
		Gender:         u.Gender,
		EnrollmentYear: u.EnrollmentYear,
		MobilePhone:    u.MobilePhone,
		College:        u.College,
		Nickname:       u.Nickname,
		Realname:       u.Realname,
		HasProfile:     u.HasProfile,
		AvatarURL:      u.AvatarUrl,
	}
	if err := db.AddOrUpdateUserProfileInfo(dbu); err != nil {
		return err
	}
	return nil
}

func (s *UploadUserService) UploadUserProfileInfo(u *user.UserProfileInfo) error {
	s.UploadUserInfo(u.UserInfo)
	dbu := &db.UserProfileInfo{
		UserID: u.UserInfo.UserId,
		//TODO:contestfavoritecount
		Introduction: u.Introduction,
		QQNumber:     u.QqNumber,
		WeChatNumber: u.WechatNumber,
	}
	if err := db.AddOrUpdateUserProfileInfo(dbu); err != nil {
		return err
	}
	if err := db.AddOrUpdateHonors(u.UserInfo.UserId, u.Honors); err != nil {
		return err
	}
	return nil
}
