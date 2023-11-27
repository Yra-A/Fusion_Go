# Fusion_Go

Fusion 是一个提供了高校赛事信息汇总以及组队功能的网站，项目采用前后端分离，前端主要采用 Vue 框架，后端主要基于 Go 的微服务框架。
该 Repo 保存的是后端代码。

前端项目地址：[Fusion_Web](https://github.com/Yra-A/Fusion_Web)

## API 接口文档
查看详细的 API 接口文档：[Fusion_API](https://apifox.com/apidoc/shared-f57b848c-5302-494c-b088-6e2443d055bb)


## 项目介绍
1. **微服务架构**: 使用 HTTP 框架 [Hertz](https://github.com/cloudwego/kitex) 构建 API 层，微服务模块之间通过 RPC 框架 [Kitex](https://github.com/cloudwego/kitex) 进行通信。

2. **数据库操作**: 采用 GORM 框架操作 MySQL 数据库，实现高效的数据处理和存储。

3. **数据缓存**: 使用 Redis 对频繁访问的数据进行缓存，显著提升接口性能。

4. **服务发现与注册**: 利用 ETCD 进行服务发现和注册，增强系统的可伸缩性和可靠性。

5. **用户认证**: 使用 JWT 实现用户 token 的生成和校验，确保系统安全性。

6. **链路跟踪**: 通过 Hertz 中间件 tracer 实现链路跟踪，保证高可维护性和快速问题定位。

7. **日志记录**: 结合 Kitex 中间件 klog 和 Hertz 中间件 hlog 进行全面的日志记录。

8. **Gzip 压缩**: 利用 gzip 中间件对 API 响应进行压缩，减少数据传输量，提高响应速度，特别适用于高延迟和低带宽的网络环境。适用于高延迟和低带宽的网络环境。

## 快速开始
1. 编辑 pkg/constants/constants.go 文件，修改相关配置

2. 启动相关服务，保证已经安装了 docker
```shell
make start
```

3. 启动 api 服务
```shell
make run_api
```

4. 启动 user 服务
```shell
make run_user
```

5. 启动 contest 服务
```shell
make run_contest
```

6. 启动 team 服务
```shell
make run_team
```

7. 启动 favorite 服务
```shell
make run_favorite
```