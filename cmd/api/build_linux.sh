#!/bin/bash
RUN_NAME=hertz_service
mkdir -p output/bin
cp script/* output 2>/dev/null
chmod +x output/bootstrap.sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o output/bin/${RUN_NAME}