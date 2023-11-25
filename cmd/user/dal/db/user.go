package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserProfileInfo struct {
	UserID         int32  `gorm:"primary_key;column:user_id"`
	Gender         int32  `gorm:"column:gender"`
	EnrollmentYear int32  `gorm:"column:enrollment_year"`
	MobilePhone    string `gorm:"column:mobile_phone"`
	College        string `gorm:"column:college"`
	Nickname       string `gorm:"column:nickname"`
	Realname       string `gorm:"column:realname"`
	AvatarURL      string `gorm:"column:avatar_url"`
	HasProfile     bool   `gorm:"column:hasProfile"`
	Introduction   string `gorm:"column:introduction"`
	QQNumber       string `gorm:"column:qq_number"`
	WeChatNumber   string `gorm:"column:wechat_number"`
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
	return DB.Transaction(func(tx *gorm.DB) error {
		u := &Authentication{
			Username: username,
			Password: password,
		}
		if err := tx.Create(u).Error; err != nil {
			return err
		}
		profile := &UserProfileInfo{UserID: u.UserID}
		if err := tx.Create(profile).Error; err != nil {
			return err
		}
		return nil
	})
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
func QueryUserProfileByUserId(tx *gorm.DB, userId int32) (*UserProfileInfo, error) {
	u := &UserProfileInfo{}
	if err := DB.Where("user_id = ?", userId).First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// AddUserProfileInfo add user profile info
func AddOrUpdateUserProfileInfo(u *UserProfileInfo) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		existingProfile, err := QueryUserProfileByUserId(tx, u.UserID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if existingProfile != nil {
			return tx.Model(existingProfile).Updates(u).Error
		}
		return tx.Create(u).Error
	})
}

// QueryHonorsByUserId 获取某个用户的所有荣誉
func QueryHonorsByUserId(userId int32) ([]string, error) {
	var honors []Honors
	if err := DB.Where("user_id = ?", userId).Find(&honors).Error; err != nil {
		return nil, err
	}
	honorStrings := make([]string, 0, len(honors))
	for _, honor := range honors {
		honorStrings = append(honorStrings, honor.Honor)
	}
	return honorStrings, nil
}

// AddOrUpdateHonors 更新用户的整个荣誉列表
func AddOrUpdateHonors(userId int32, honors []string) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		// 删除该用户的所有现有荣誉
		if err := tx.Where("user_id = ?", userId).Delete(&Honors{}).Error; err != nil {
			return err
		}
		for _, honor := range honors {
			newHonor := Honors{
				UserID: userId,
				Honor:  honor,
			}
			if err := tx.Create(&newHonor).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// UpdateHasProfile 更新用户是否填写了个人档案的状态
func UpdateHasProfile(userId int32, hasProfile bool) error {
	return DB.Model(&UserProfileInfo{}).Where("user_id = ?", userId).Update("hasProfile", hasProfile).Error
}
