package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserProfileInfo struct {
	UserID               int32  `gorm:"primary_key;column:user_id"`
	Gender               int32  `gorm:"column:gender"`
	EnrollmentYear       int32  `gorm:"column:enrollment_year"`
	MobilePhone          string `gorm:"column:mobile_phone"`
	College              string `gorm:"column:college"`
	Nickname             string `gorm:"column:nickname"`
	Realname             string `gorm:"column:realname"`
	AvatarURL            string `gorm:"column:avatar_url"`
	HasProfile           bool   `gorm:"column:hasProfile"`
	ContestFavoriteCount int32  `gorm:"column:contest_favorite_count"`
	Introduction         string `gorm:"column:introduction"`
	QQNumber             string `gorm:"column:qq_number"`
	WeChatNumber         string `gorm:"column:wechat_number"`
}

func (UserProfileInfo) TableName() string {
	return "user_profile_info"
}

type UserInfo struct {
	UserID         int32
	Gender         int32
	EnrollmentYear int32
	MobilePhone    string
	College        string
	Nickname       string
	Realname       string
	AvatarURL      string
	HasProfile     bool
}

type UserProfile struct {
	UserID               int32
	ContestFavoriteCount int32
	Introduction         string
	QQNumber             string
	WeChatNumber         string
}

type Honors struct {
	HonorID int32  `gorm:"primary_key;column:honor_id"`
	UserID  int32  `gorm:"column:user_id"`
	Honor   string `gorm:"column:honor"`
}

func (Honors) TableName() string {
	return "honors"
}

type Authentication struct {
	UserID   int32  `gorm:"primary_key;autoIncrement;column:user_id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (Authentication) TableName() string {
	return "authentication"
}

// QueryUserByName query user by username
func QueryUserByName(username string) (*Authentication, error) {
	u := &Authentication{}
	if err := DB.Where("username = ?", username).First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func CreateUser(username, password string) error {
	u := &Authentication{
		Username: username,
		Password: password,
	}
	if err := DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

// QueryUserByUserId query user by user_id
func QueryUserByUserId(userId int32) (*UserInfo, error) {
	var profile UserProfileInfo
	if err := DB.Where("user_id = ?", userId).First(&profile).Error; err != nil {
		return nil, err
	}
	userInfo := &UserInfo{
		UserID:         profile.UserID,
		Gender:         profile.Gender,
		EnrollmentYear: profile.EnrollmentYear,
		MobilePhone:    profile.MobilePhone,
		College:        profile.College,
		Nickname:       profile.Nickname,
		Realname:       profile.Realname,
		AvatarURL:      profile.AvatarURL,
		HasProfile:     profile.HasProfile,
	}

	return userInfo, nil
}

// QueryUserProfileByUserId query user profile by user_id
func QueryUserProfileByUserId(userId int32) (*UserProfileInfo, error) {
	u := &UserProfileInfo{}
	if err := DB.Where("user_id = ?", userId).First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// AddUserProfileInfo add user profile info
func AddOrUpdateUserProfileInfo(u *UserProfileInfo) error {
	existingProfile, err := QueryUserProfileByUserId(u.UserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existingProfile != nil {
		return DB.Model(existingProfile).Updates(u).Error
	}
	return DB.Create(u).Error
}

// QueryHonorsByUserId 获取某个用户的所有荣誉
func QueryHonorsByUserId(userId int32) ([]string, error) {
	var honors []string
	if err := DB.Where("user_id = ?", userId).Find(&honors).Error; err != nil {
		return nil, err
	}
	return honors, nil

}
