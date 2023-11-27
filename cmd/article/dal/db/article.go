package db

import (
	"time"
)

type Article struct {
	ArticleID   int32     `gorm:"primary_key;column:article_id"`
	Title       string    `gorm:"column:title;not null"`
	AuthorID    int32     `gorm:"column:author_id"`
	Author      string    `gorm:"column:author"`
	Link        string    `gorm:"column:link;not null"`
	CreatedTime time.Time `gorm:"column:created_time"`
	ContestID   int32     `gorm:"column:contest_id;not null"`
}

func (Article) TableName() string {
	return "article"
}

// FetchArticleList 根据contest_id, limit, offset来获取文章列表
func FetchArticleList(contestId int32, limit int32, offset int32) ([]*Article, error) {
	var articleBriefInfos []*Article

	//将*gorm.DB实例与Article模型关联，并能倒序
	query := DB.Model(&Article{}).Order("created_time desc")

	//筛选
	query = query.Where("contest_id = ?", contestId)

	//执行select语句，确保字段名和ArticleBrief结构中一致
	query = query.Select("article_id, title, author_id, author, created_time, link")

	//应用分页
	query = query.Offset(int(offset)).Limit(int(limit))

	//执行查询
	err := query.Find(&articleBriefInfos).Error
	if err != nil {
		return nil, err
	}

	return articleBriefInfos, nil
}
