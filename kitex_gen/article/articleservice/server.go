// Code generated by Kitex v0.7.3. DO NOT EDIT.
package articleservice

import (
	article "github.com/Yra-A/Fusion_Go/kitex_gen/article"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler article.ArticleService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
