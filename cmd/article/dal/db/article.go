package db

import (
	"github.com/Yra-A/Fusion_Go/kitex_gen/article"
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

// mock
/*var articlesData = map[int32]*Article{
	1: {
		ArticleID:   1,
		Title:       "369酒桶玩得跟屎一样",
		AuthorID:    1,
		Author:      "上海师范大学高振宁老师",
		Link:        "https://www.bilibili.com/video/BV15w411T7Xc",
		CreatedTime: 1610000000,
		ContestID:   1,
	},
	2: {
		ArticleID:   2,
		Title:       "决赛燃烧自己",
		AuthorID:    2,
		Author:      "上海师范大学李元浩老师",
		Link:        "https://www.bilibili.com/video/BV1Ju411F7Jt",
		CreatedTime: 1620000000,
		ContestID:   1,
	},
	3: {
		ArticleID:   3,
		Title:       "红尘客栈",
		AuthorID:    3,
		Author:      "上海师范大学周杰伦老师",
		Link:        "https://www.bilibili.com/video/BV1FT411m7d6",
		CreatedTime: 1630000000,
		ContestID:   2,
	},
	4: {
		ArticleID:   4,
		Title:       "雨下一整晚",
		AuthorID:    3,
		Author:      "上海师范大学周杰伦老师",
		Link:        "https://www.bilibili.com/video/BV1EM41117BM",
		CreatedTime: 1640000000,
		ContestID:   2,
	},
	//其他Article数据
}
*/

// FetchArticleList 根据contest_id, limit, offset来获取文章列表
func FetchArticleList(contestId int32, limit int32, offset int32) ([]*article.ArticleBrief, error) {
	var articleBriefInfos []*article.ArticleBrief

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

/*
// FetchArticleListMock is a mock function to simulate database behavior for testing purposes.
func FetchArticleListMock(contestId int32, limit int32, offset int32) ([]*article.ArticleBrief, error) {
	// 将数据库中符合条件的数据放到articlesSlice中
	articlesSlice := make([]*Article, 0, len(articlesData))
	for _, v := range articlesData {
		articlesSlice = append(articlesSlice, v)
	}

	//按CreatedTime降序排列
	sort.Slice(articlesSlice, func(i, j int) bool {
		return articlesSlice[i].CreatedTime > articlesSlice[j].CreatedTime
	})

	//根据contest_id过滤排序后的articlesSlice
	var filteredSortedArticles []*Article
	for _, v := range articlesSlice {
		if v.ContestID == contestId {
			filteredSortedArticles = append(filteredSortedArticles, v)
		}
	}

	//应用 offset 和 limit
	start := int(offset)
	if start >= len(filteredSortedArticles) {
		return nil, nil
	}
	end := start + int(limit)
	if end > len(filteredSortedArticles) {
		end = len(filteredSortedArticles)
	}
	paginatedArticles := filteredSortedArticles[start:end]

	//创建 ArticleBriefInfo 列表
	briefInfos := make([]*article.ArticleBrief, 0, len(paginatedArticles))

	for _, v := range paginatedArticles {
		briefInfo := &article.ArticleBrief{
			ArticleId:   v.ArticleID,
			Title:       v.Title,
			AuthorId:    v.AuthorID,
			Author:      v.Author,
			CreatedTime: v.CreatedTime,
			Link:        v.Link,
		}
		briefInfos = append(briefInfos, briefInfo)
	}
	return briefInfos, nil
}
*/
