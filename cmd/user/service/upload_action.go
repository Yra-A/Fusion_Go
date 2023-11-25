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
	if err := s.UploadUserInfo(u.UserInfo); err != nil {
		return err
	}
	dbu := &db.UserProfileInfo{
		UserID:       u.UserInfo.UserId,
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
	return db.UpdateHasProfile(u.UserInfo.UserId, true)
}

func (s *UploadUserService) UpdateHasProfile(userId int32) error {
	profile, err := db.QueryUserProfileByUserId(db.DB, userId)
	if err != nil {
		return err
	}

	honors, err := db.QueryHonorsByUserId(userId)
	if err != nil {
		return err
	}

	// 检查至少有一个字段是否有值
	hasProfile := profile.Introduction != "" || profile.QQNumber != "" || profile.WeChatNumber != "" || len(honors) > 0
	profile.HasProfile = hasProfile

	// 更新 UserProfile 信息
	return db.UpdateHasProfile(userId, hasProfile)
}
