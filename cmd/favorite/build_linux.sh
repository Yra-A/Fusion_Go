#!/usr/bin/env bash
RUN_NAME="FavoriteService"

mkdir -p output/bin
cp script/* output/
chmod +x output/bootstrap.sh

if [ "$IS_SYSTEM_TEST_ENV" != "1" ]; then
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o output/bin/${RUN_NAME}
else
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test -c -covermode=set -o output/bin/${RUN_NAME} -coverpkg=./...
fi
