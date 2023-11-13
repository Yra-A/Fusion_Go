## 启动 api 层
run_api:
	cd cmd/api && sh ./build.sh && sh ./output/bootstrap.sh

## 启动 user 服务
run_user:
	cd cmd/user && sh ./build.sh && sh ./output/bootstrap.sh

## 启动 team 服务
run_team:
	cd cmd/team && sh ./build.sh && sh ./output/bootstrap.sh

## 启动相关服务
start:
	docker compose --profile dev up -d

## 关闭相关服务
stop:
	docker-compose --profile dev stop

## 关闭并删除
down:
	docker-compose --profile dev down