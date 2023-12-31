// Code generated by hertz generator.

package main

import (
	"github.com/Yra-A/Fusion_Go/cmd/api/biz/mw/jwt"
	"github.com/Yra-A/Fusion_Go/cmd/api/biz/mw/oss"
	"github.com/Yra-A/Fusion_Go/cmd/api/rpc"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"time"
)

func Init() {
	rpc.InitRPC()
	jwt.InitJwt()
	oss.Init()
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(hlog.LevelInfo)
}

func main() {
	Init()
	tracer, cfg := tracing.NewServerTracer()
	h := server.New(
		server.WithStreamBody(true),
		server.WithHostPorts("0.0.0.0:8888"),
		tracer,
	)
	h.Use(accesslog.New(accesslog.WithFormat("[${url}=-=-=${time}] ${status} - ${latency} ${method} ${path} ${queryParams} - 【req body: ${body}】【req query parameter: ${queryParams}】【response body: ${resBody}】")))
	h.Use(tracing.ServerMiddleware(cfg))
	h.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // 允许来自任意 origin 的客户端访问服务端资源
		//AllowOrigins: []string{"http://101.35.229.143", "http://101.35.229.143/"}, // 允许添加的请求源
		AllowMethods: []string{"POST", "GET", "OPTIONS"}, // 允许添加的请求方法
		//AllowHeaders:     []string{"Content-Type", "Authorization", "Origin", "Content-Length"}, // 允许添加的请求头
		AllowHeaders:     []string{"*"},              // 允许添加的请求头
		ExposeHeaders:    []string{"Content-Length"}, // 暴露给客户端的响应头
		AllowCredentials: true,                       // 前端请求携带凭证如Cookies或HTTP认证的时候需要设置为true
		MaxAge:           12 * time.Hour,
	}))
	register(h)
	h.Spin()
}
