package db

import (
	"errors"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
	"gorm.io/gorm"
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

// CreateArticle 创建文章
func CreateArticle(title string, authorId int32, author string, link string, contestId int32) (int32, error) {
	article := &Article{
		Title:       title,
		AuthorID:    authorId,
		Author:      author,
		Link:        link,
		CreatedTime: time.Now(),
		ContestID:   contestId,
	}
	err := DB.Create(article).Error
	if err != nil {
		return 0, err
	}
	return article.ArticleID, nil
}

func ModifyArticle(articleId int32, title string, authorId int32, author string, link string, contestId int32) (int32, error) {
	article := &Article{}
	err := DB.Model(&Article{}).Where("article_id = ?", articleId).First(&article).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errno.ArticleNotExistErr
	}
	if err != nil {
		return 0, err
	}

	articles := &Article{
		ArticleID: articleId,
	}
	err = DB.Model(&articles).Updates(map[string]interface{}{
		"title":      title,
		"author_id":  authorId,
		"author":     author,
		"link":       link,
		"contest_id": contestId,
	}).Error

	if err != nil {
		return 0, err
	}
	return article.ArticleID, nil
}

// FetchArticleList 根据contest_id, limit, offset来获取文章列表
func FetchArticleList(contestId int32, limit int32, offset int32) ([]*Article, int32, error) {
	var articleBriefInfos []*Article

	//将*gorm.DB实例与Article模型关联，并能倒序
	query := DB.Model(&Article{}).Order("created_time desc")

	//筛选
	query = query.Where("contest_id = ?", contestId)

	//执行select语句，确保字段名和ArticleBrief结构中一致
	query = query.Select("article_id, title, author_id, author, created_time, link")

	var total int64

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	//应用分页
	query = query.Offset(int(offset)).Limit(int(limit))

	//执行查询
	err := query.Find(&articleBriefInfos).Error
	if err != nil {
		return nil, 0, err
	}

	return articleBriefInfos, int32(total), nil
}
