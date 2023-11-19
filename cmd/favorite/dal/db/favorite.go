package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

// UserFavorite corresponds to the 'user_favorites' table in the database.
type UserFavorite struct {
	FavoriteID  int32     `gorm:"primary_key;auto_increment;column:favorite_id"`
	UserID      int32     `gorm:"column:user_id"`
	ContestID   int32     `gorm:"column:contest_id"`
	CreatedTime time.Time `gorm:"column:created_time"`
}

func (UserFavorite) TableName() string {
	return "user_favorites"
}

func ContestFavoriteAction(user_id int32, contest_id int32, action_type int32) error {
	var userFavorite UserFavorite
	err := DB.Where("user_id = ? AND contest_id = ?", user_id, contest_id).First(&userFavorite).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if action_type == 1 {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return DB.Create(&UserFavorite{
				UserID:      user_id,
				ContestID:   contest_id,
				CreatedTime: time.Now(),
			}).Error
		}
		return errors.New("already favorited")
	} else if action_type == 2 {
		if err == nil {
			return DB.Delete(&userFavorite).Error
		}
		return errors.New("not favorited yet")
	}
	return errors.New("invalid action type")
}

func FetchContestFavoriteList(user_id int32, limit int32, offset int32) ([]int32, error) {
	var userFavorites []*UserFavorite
	err := DB.Where("user_id = ?", user_id).Order("created_time DESC").Limit(int(limit)).Offset(int(offset)).Find(&userFavorites).Error
	if err != nil {
		return nil, err
	}
	contestIds := make([]int32, len(userFavorites))
	for i, v := range userFavorites {
		contestIds[i] = v.ContestID
	}
	return contestIds, nil
}

func QueryFavoriteStatusByUserId(user_id int32, contest_id int32) (bool, error) {
	var userFavorite UserFavorite
	err := DB.Where("user_id = ? AND contest_id = ?", user_id, contest_id).First(&userFavorite).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, nil
}
