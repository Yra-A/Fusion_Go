package db

import (
	"github.com/Yra-A/Fusion_Go/cmd/favorite/dal/redis"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

var rdFav redis.Favorite

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

	// 出现了 找不到记录之外的错误
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if action_type == 1 {
		// 没有找到该条记录
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 更新 db，创建一条收藏记录
			u := &UserFavorite{
				UserID:      user_id,
				ContestID:   contest_id,
				CreatedTime: time.Now(),
			}
			err := DB.Create(u).Error
			// 异步更新 cache
			go func() {
				rdFav.AddFavoriteContest(int64(user_id), int64(contest_id), u.CreatedTime)
			}()
			return err
		}
		return errors.New("already favorited")
	} else if action_type == 2 {
		// 找到了该条记录进行删除
		if err == nil {
			// 更新 db
			err := DB.Delete(&userFavorite).Error
			// 异步更新 cache
			go func() {
				rdFav.DelFavoriteContest(int64(user_id), int64(contest_id))
			}()
			return err
		}
		return errors.New("not favorited yet")
	}
	return errors.New("invalid action type")
}

func FetchContestFavoriteList(user_id int32, limit int32, offset int32) ([]int32, error) {
	// 先查询 cache，如果有则直接返回
	if rdFav.CheckFavoriteContest(int64(user_id)) {
		contestId_list_i64 := rdFav.GetFavoriteContest(int64(user_id))
		contestId_list_i32 := make([]int32, len(contestId_list_i64))
		for i := range contestId_list_i64 {
			contestId_list_i32[i] = int32(contestId_list_i64[i])
		}
		return contestId_list_i32, nil
	}
	// cache 中没有记录，则先查询数据库，随后异步更新到缓存中
	var userFavorites []*UserFavorite
	err := DB.Where("user_id = ?", user_id).Order("created_time DESC").Limit(int(limit)).Offset(int(offset)).Find(&userFavorites).Error
	if err != nil {
		return nil, err
	}
	contestIds := make([]int32, len(userFavorites))
	for i, v := range userFavorites {
		contestIds[i] = v.ContestID
	}
	// 将从 db 中获取到的 contest id list 异步更新到 cache 中
	go func(userFavorites []*UserFavorite) {
		for _, fav := range userFavorites {
			rdFav.AddFavoriteContest(int64(fav.UserID), int64(fav.ContestID), fav.CreatedTime)
		}
	}(userFavorites)
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
