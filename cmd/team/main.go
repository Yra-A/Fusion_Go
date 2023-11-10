// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
    "github.com/Yra-A/Fusion_Go/cmd/team/dal"
    "github.com/Yra-A/Fusion_Go/cmd/team/rpc"
    "github.com/kitex-contrib/obs-opentelemetry/tracing"
    "net"

    team "github.com/Yra-A/Fusion_Go/kitex_gen/team/teamservice"
    "github.com/Yra-A/Fusion_Go/pkg/constants"
    "github.com/Yra-A/Fusion_Go/pkg/middleware"
    "github.com/cloudwego/kitex/pkg/klog"
    "github.com/cloudwego/kitex/pkg/limit"
    "github.com/cloudwego/kitex/pkg/rpcinfo"
    "github.com/cloudwego/kitex/server"
    kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
    etcd "github.com/kitex-contrib/registry-etcd"
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

    addr, err := net.ResolveTCPAddr("tcp", "localhost:8895")
    if err != nil {
        panic(err)
    }
    Init()

    svr := team.NewServer(new(TeamServiceImpl),
        server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.TeamServiceName}), // server name
        server.WithMiddleware(middleware.CommonMiddleware),                                             // middleware
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