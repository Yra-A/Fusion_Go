package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/article/dal/db"
	"github.com/Yra-A/Fusion_Go/kitex_gen/article"
)

// QueryArticleListService “用于查询文章列表的服务”类型
type QueryArticleListService struct {
	ctx context.Context
}

// QueryArticleList “查询比赛列表”方法
func (s *QueryArticleListService) QueryArticleList(contestId int32, limit int32, offset int32) ([]*article.ArticleBriefInfo, int32, error) {
	dbArticles, total, err := db.FetchArticleList(contestId, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	//将dbArticles(db.Article类型)转换为articleBriefInfo(ArticleBriefInfo类型)
	articleBriefInfos := make([]*article.ArticleBriefInfo, len(dbArticles))

	for i, v := range dbArticles {
		briefInfo := &article.ArticleBriefInfo{
			ArticleBriefInfo: &article.ArticleBrief{
				ArticleId:   v.ArticleID,
				Title:       v.Title,
				AuthorId:    v.AuthorID,
				Author:      v.Author,
				CreatedTime: v.CreatedTime.Unix(),
				Link:        v.Link,
			},
		}
		articleBriefInfos[i] = briefInfo
	}

	return articleBriefInfos, total, nil
}

// NewQueryArticleListService 构造函数，用于创建实例，传入上下文
func NewQueryArticleListService(ctx context.Context) *QueryArticleListService {
	return &QueryArticleListService{ctx: ctx}
}
