package service

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/article/dal/db"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
)

type CreateArticleService struct {
	ctx context.Context
}

func (s *CreateArticleService) CreateArticle(articleId int32, title string, authorId int32, author string, link string, contestId int32) (int32, error) {
	if articleId == 0 {
		return db.CreateArticle(title, authorId, author, link, contestId)
	} else if articleId > 0 {
		return db.ModifyArticle(articleId, title, authorId, author, link, contestId)
	}

	return 0, errno.ParamErr
}

func NewCreateArticleService(ctx context.Context) *CreateArticleService {
	return &CreateArticleService{ctx: ctx}
}
