package main

import (
	"context"
	"github.com/Yra-A/Fusion_Go/cmd/article/service"
	article "github.com/Yra-A/Fusion_Go/kitex_gen/article"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

// ArticleServiceImpl implements the last service interface defined in the IDL.
type ArticleServiceImpl struct{}

// ArticleList implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) ArticleList(ctx context.Context, req *article.ArticleListRequest) (resp *article.ArticleListResponse, err error) {
	klog.CtxDebugf(ctx, "ArticleList called")
	resp = new(article.ArticleListResponse)
	c, err := service.NewQueryArticleListService(ctx).QueryArticleList(req.ContestId, req.Limit, req.Offset)

	//返回查询过程中可能的错误信息
	if err != nil {
		resp.StatusCode = errno.Fail.ErrCode
		resp.StatusMsg = errno.Fail.ErrMsg
		return resp, nil
	}

	//返回查询成功的信息
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	resp.Total = int32(len(c))
	resp.ArticleList = c
	return resp, nil
}
