package rpc

import (
	"context"
	"github.com/Yra-A/Fusion_Go/kitex_gen/article"
	"github.com/Yra-A/Fusion_Go/kitex_gen/article/articleservice"
	"github.com/Yra-A/Fusion_Go/pkg/constants"
	"github.com/Yra-A/Fusion_Go/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var articleClient articleservice.Client

func initArticleRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := articleservice.NewClient(
		constants.ArticleServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	articleClient = c //articleClient就是新rpc客户端
}

// ArticleCreate 创建文章【rpc 客户端】
func ArticleCreate(ctx context.Context, req *article.ArticleCreateRequest) (*article.ArticleCreateResponse, error) {
	resp, err := articleClient.ArticleCreate(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ArticleList 文章列表【rpc 客户端】
func ArticleList(ctx context.Context, req *article.ArticleListRequest) (*article.ArticleListResponse, error) {
	resp, err := articleClient.ArticleList(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
