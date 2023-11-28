## 启动 api 层
run_api:
	cd cmd/api && sh ./build.sh && sh ./output/bootstrap.sh

## 启动 user 服务
run_user:
	cd cmd/user && sh ./build.sh && sh ./output/bootstrap.sh

## 启动 contest 服务
run_contest:
	cd cmd/contest && sh ./build.sh && sh ./output/bootstrap.sh

## 启动 article 服务
run_article:
	cd cmd/article && sh ./build.sh && sh ./output/bootstrap.sh

## 启动 team 服务
run_team:
	cd cmd/team && sh ./build.sh && sh ./output/bootstrap.sh

## 启动 favorite 服务
run_favorite:
	cd cmd/favorite && sh ./build.sh && sh ./output/bootstrap.sh

## 启动相关服务
start:
	docker compose --profile dev up -d

## 关闭相关服务
stop:
	docker-compose --profile dev stop

## 关闭并删除
down:
	docker-compose --profile dev down

## 打包
build:
	rm -rf Fusion_Go_build
	mkdir -p Fusion_Go_build/api_output
	mkdir -p Fusion_Go_build/user_output
	mkdir -p Fusion_Go_build/favorite_output
	mkdir -p Fusion_Go_build/contest_output
	mkdir -p Fusion_Go_build/team_output
	cd cmd/api && sh ./build_linux.sh && cp -r output/* ../../Fusion_Go_build/api_output
	cd cmd/user && sh ./build_linux.sh && cp -r output/* ../../Fusion_Go_build/user_output
	cd cmd/favorite && sh ./build_linux.sh && cp -r output/* ../../Fusion_Go_build/favorite_output
	cd cmd/contest && sh ./build_linux.sh && cp -r output/* ../../Fusion_Go_build/contest_output
	cd cmd/team && sh ./build_linux.sh && cp -r output/* ../../Fusion_Go_build/team_output
	cp Makefile_Fusion Fusion_Go_build/Makefile