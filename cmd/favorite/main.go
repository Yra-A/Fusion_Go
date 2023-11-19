package main

import (
	"github.com/Yra-A/Fusion_Go/cmd/favorite/dal"
	"github.com/Yra-A/Fusion_Go/cmd/favorite/rpc"
	favorite "github.com/Yra-A/Fusion_Go/kitex_gen/favorite/favoriteservice"
	"github.com/Yra-A/Fusion_Go/pkg/constants"
	"github.com/Yra-A/Fusion_Go/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

func Init() {
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	dal.Init()
	rpc.InitRPC()
}
func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8894")
	if err != nil {
		panic(err)
	}
	Init()
	svr := favorite.NewServer(new(FavoriteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.FavoriteServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                                 // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(tracing.NewServerSuite()),                         // tracer
		//server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r), // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
