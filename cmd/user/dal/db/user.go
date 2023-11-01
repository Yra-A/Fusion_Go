package db

import (
	"fmt"
	"gorm.io/gorm"
)

// ===============User================
type UserInfo struct {
	gorm.Model
	UserID         int32  `gorm:"primary_key;column:user_id"`
	Gender         int32  `gorm:"column:gender"`
	EnrollmentYear int32  `gorm:"column:enrollment_year"`
	MobilePhone    string `gorm:"column:mobile_phone"`
	College        string `gorm:"column:college"`
	Nickname       string `gorm:"column:nickname"`
	Realname       string `gorm:"column:realname"`
	AvatarURL      string `gorm:"column:avatar_url"`
	HasProfile     bool   `gorm:"column:hasProfile"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

type UserProfileInfo struct {
	gorm.Model
	UserID               int32  `gorm:"primary_key;column:user_id"`
	ContestFavoriteCount int32  `gorm:"column:contest_favorite_count"`
	Introduction         string `gorm:"column:introduction"`
	QQNumber             string `gorm:"column:qq_number"`
	WeChatNumber         string `gorm:"column:wechat_number"`
}

func (UserProfileInfo) TableName() string {
	return "user_profile_info"
}

type Honors struct {
	gorm.Model
	HonorID int32  `gorm:"primary_key;column:honor_id"`
	UserID  int32  `gorm:"column:user_id"`
	Honor   string `gorm:"column:honor"`
}

func (Honors) TableName() string {
	return "honors"
}

type Image struct {
	gorm.Model
	ImageID  int32  `gorm:"primary_key;column:image_id"`
	UserID   int32  `gorm:"column:user_id"`
	ImageURL string `gorm:"column:image_url"`
}

func (Image) TableName() string {
	return "images"
}

type Authentication struct {
	AuthID   int32  `gorm:"primaryKey;autoIncrement;column:auth_id"`
	UserID   int32  `gorm:"column:user_id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

var (
	userInfo = &UserInfo{
		UserID:         1,
		Gender:         1,
		EnrollmentYear: 2023,
		MobilePhone:    "1234567890",
		College:        "Example College",
		Nickname:       "nickname_example",
		Realname:       "realname_example",
		AvatarURL:      "http://example.com/avatar.jpg",
		HasProfile:     true,
	}

	userProfileInfo = &UserProfileInfo{
		UserID:               1,
		ContestFavoriteCount: 5,
		Introduction:         "This is an example introduction.",
		QQNumber:             "12345678",
		WeChatNumber:         "wechat_example",
	}

	authentication = &Authentication{
		AuthID:   1,
		UserID:   1,
		Username: "rocketzhu",
		Password: "123456",
	}
)

// Mock data for images
var mockHonorsData = []Honors{
	{
		HonorID: 1,
		UserID:  3,
		Honor:   "Honor 1",
	},
	{
		HonorID: 2,
		UserID:  3,
		Honor:   "Honor 2",
	},
}

// Mock data for images
var mockImagesData = []Image{
	{
		ImageID:  1,
		UserID:   3,
		ImageURL: "http://example.com/image1.jpg",
	},
	{
		ImageID:  2,
		UserID:   3,
		ImageURL: "http://example.com/image2.jpg",
	},
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

// QueryUserByName query user by username
func QueryUserByName(username string) (*Authentication, error) {
	//u := &Authentication{}
	//if err := DB.Where("username = ?", username).First(u).Error; err != nil {
	//	return nil, err
	//}
	u := authentication
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

// QueryHonorsByUserId 获取某个用户的所有荣誉
func QueryHonorsByUserId(userId int32) ([]string, error) {
	var honors []Honors
	//if err := DB.Where("user_id = ?", userId).Find(&honors).Error; err != nil {
	//	return nil, err
	//}
	honors = mockHonorsData
	honorDescriptions := make([]string, len(honors))
	for i := 0; i < len(honors); i++ {
		honorDescriptions[i] = honors[i].Honor
	}
	return honorDescriptions, nil
}

// QuseryImagesByUserId 获取某个用户的所有图片
func QueryImagesByUserId(userId int32) ([]string, error) {
	var images []Image
	//if err := DB.Where("user_id = ?", userId).Find(&images).Error; err != nil {
	//	return nil, err
	//}
	images = mockImagesData
	imageURLs := make([]string, len(images))
	for i := 0; i < len(images); i++ {
		imageURLs[i] = images[i].ImageURL
	}
	return imageURLs, nil
}
