package redis

import (
	"strconv"
	"time"
)

const (
	userIdPrefix       = "user_id:"
	favorContestSuffix = "_favor_contest"
)

type (
	Favorite struct{}
)

func getFavoriteContestKeyStr(user_id int64) string {
	return userIdPrefix + strconv.FormatInt(user_id, 10) + favorContestSuffix
}

// AddFavoriteContest 往集合 user_id:xxx_favor_contest 中添加 contest_id
func (f Favorite) AddFavoriteContest(user_id, contest_id int64, createdTime time.Time) {
	// 使用负的Unix时间戳作为分数,实现按照时间降序排序
	score := float64(-createdTime.Unix())
	add(rdb, getFavoriteContestKeyStr(user_id), contest_id, score)
}

// DelFavoriteContest 往集合 user_id:xxx_favor_contest 中添加 contest_id
func (f Favorite) DelFavoriteContest(user_id, contest_id int64) {
	del(rdb, getFavoriteContestKeyStr(user_id), contest_id)
}

// CheckFavoriteContest 检查集合 user_id:xxx_favor_contest 是否存在
func (f Favorite) CheckFavoriteContest(user_id int64) bool {
	return check(rdb, getFavoriteContestKeyStr(user_id))
}

// ExistFavoriteContest 检查集合 user_id:xxx_favor_contest 中是否存在 contest_id
func (f Favorite) ExistFavoriteContest(user_id, contest_id int64) bool {
	return exist(rdb, getFavoriteContestKeyStr(user_id), contest_id)
}

// CountFavoriteContest 统计集合 user_id:xxx_favor_contest 中的数量
func (f Favorite) CountFavoriteContest(user_id int64) (int64, error) {
	return count(rdb, getFavoriteContestKeyStr(user_id))
}

// GetFavoriteContest 获取集合 user_id:xxx_favor_contest 中的所有 contest_id
func (f Favorite) GetFavoriteContest(user_id int64) []int64 {
	return get(rdb, getFavoriteContestKeyStr(user_id))
}
