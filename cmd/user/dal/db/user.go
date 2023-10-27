package db

import (
	"fmt"
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	UserId               int32  `gorm:"primaryKey" json:"user_id"`
	Gender               int32  `gorm:"" json:"gender"`
	Nickname             string `gorm:"type:varchar(255)" json:"nickname"`
	Realname             string `gorm:"type:varchar(255)" json:"realname"`
	ContestFavoriteCount int32  `gorm:"" json:"contest_favorite_count"`
	AvatarUrl            string `gorm:"type:varchar(255)" json:"avatar_url"`
	EnrollmentYear       int32  `gorm:"" json:"enrollment_year"`
	College              string `gorm:"type:varchar(255)" json:"college"`
}

type UserProfileInfo struct {
	gorm.Model
	UserId       int32    `gorm:"primaryKey" json:"user_id"`
	MobilePhone  string   `gorm:"type:varchar(255)" json:"mobile_phone"`
	Introduction string   `gorm:"type:varchar(255)" json:"introduction"`
	QqNumber     string   `gorm:"type:varchar(255)" json:"qq_number"`
	WechatNumber string   `gorm:"type:varchar(255)" json:"wechat_number"`
	Honors       []string `gorm:"type:varchar(255)" json:"honors"`
	Images       []string `gorm:"type:varchar(255)" json:"images"`
	IsShow       []bool   `gorm:"type:varchar(255)" json:"is_show"`
}

var userInfo = &UserInfo{
	UserId:               1,
	Gender:               1,
	Nickname:             "RocketZhu",
	Realname:             "Yra-A",
	ContestFavoriteCount: 1,
	AvatarUrl:            "https://sf1-ttcdn-tos.pstatp.com/img/user-avatar/1f1a1a1a1a1a1a1a1a1a1~300x300.image",
	EnrollmentYear:       2020,
	College:              "College of Computer Science and Technology",
}

var userProfileInfo = &UserProfileInfo{
	UserId:       1,
	MobilePhone:  "12345678910",
	Introduction: "I am a student from SCUT",
	QqNumber:     "123456789",
	WechatNumber: "123456789",
	Honors:       []string{"honor1", "honor2"},
	Images:       []string{"image1", "image2"},
	IsShow:       []bool{true, false},
}

// QueryUserByUserId query user by user_id
func QueryUserByUserId(userId int32) (*UserInfo, error) {
	//u := &UserInfo{}
	//if err := DB.Where("user_id = ?", user_id).First(u).Error; err != nil {
	//	return nil, err
	//}
	fmt.Println(userId)
	u := userInfo
	return u, nil
}

// QueryUserProfileByUserId query user profile by user_id
func QueryUserProfileByUserId(userId int32) (*UserProfileInfo, error) {
	//u := &UserProfileInfo{}
	//if err := DB.Where("user_id = ?", user_id).First(u).Error; err != nil {
	//	return nil, err
	//}
	u := userProfileInfo
	return u, nil
}

// AddUserInfo add user info
func AddUserInfo(u *UserInfo) error {
	//if err := DB.Create(u).Error; err != nil {
	//	return err
	//}
	//return nil
	fmt.Printf("%#v\n", *u)
	return nil
}

// AddUserProfileInfo add user profile info
func AddUserProfileInfo(u *UserProfileInfo) error {
	//if err := DB.Create(u).Error; err != nil {
	//	return err
	//}
	//return nil
	fmt.Printf("%#v\n", *u)
	return nil
}
