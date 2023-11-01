package handler

import (
	"github.com/Yra-A/Fusion_Go/cmd/api/biz/model/api"
	"github.com/Yra-A/Fusion_Go/kitex_gen/user"
)

func ConvertUserToAPI(src *user.UserInfo) *api.UserInfo {
	return &api.UserInfo{
		UserID:         src.UserId,
		Gender:         src.Gender,
		EnrollmentYear: src.EnrollmentYear,
		MobilePhone:    src.MobilePhone,
		College:        src.College,
		Nickname:       src.Nickname,
		Realname:       src.Realname,
		HasProfile:     src.HasProfile,
		AvatarURL:      src.AvatarUrl,
	}
}

func ConvertUserProfileToAPI(src *user.UserProfileInfo) *api.UserProfileInfo {
	return &api.UserProfileInfo{
		Introduction: src.Introduction,
		QqNumber:     src.QqNumber,
		WechatNumber: src.WechatNumber,
		Honors:       src.Honors,
		Images:       src.Images,
		UserInfo:     ConvertUserToAPI(src.UserInfo),
	}
}

func ConvertAPIToUser(src *api.UserInfo) *user.UserInfo {
	return &user.UserInfo{
		UserId:         src.UserID,
		Gender:         src.Gender,
		EnrollmentYear: src.EnrollmentYear,
		MobilePhone:    src.MobilePhone,
		College:        src.College,
		Nickname:       src.Nickname,
		Realname:       src.Realname,
		HasProfile:     src.HasProfile,
		AvatarUrl:      src.AvatarURL,
	}
}

func ConvertAPIProfileToUser(src *api.UserProfileInfo) *user.UserProfileInfo {
	return &user.UserProfileInfo{
		Introduction: src.Introduction,
		QqNumber:     src.QqNumber,
		WechatNumber: src.WechatNumber,
		Honors:       src.Honors,
		Images:       src.Images,
		UserInfo:     ConvertAPIToUser(src.UserInfo),
	}
}
