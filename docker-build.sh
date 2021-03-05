#!/bin/bash

# change the image tag before building in produce for eample v0.*.*, latest version builded for testing.

APP_VERSION=latest
APP_NAME=tools

docker build -t ${APP_NAME}:${APP_VERSION} .

#docker login --username=119098598@1888046525536031.onaliyun.com registry.cn-hangzhou.aliyuncs.com

docker tag ${APP_NAME}:${APP_VERSION} registry.cn-hangzhou.aliyuncs.com/zgang/${APP_NAME}:${APP_VERSION}

docker push registry.cn-hangzhou.aliyuncs.com/zgang/${APP_NAME}:${APP_VERSION}

echo "✅ test this app:"
echo "docker run -p 8080:8080 -d ${APP_NAME}:${APP_VERSION}"

# upload remote test server
#./docker-expect.sh
