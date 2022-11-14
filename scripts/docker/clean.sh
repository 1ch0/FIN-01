#!/usr/bin/env bash

# 停止docker
docker stop $(docker ps -a | grep "Exited" | awk '{print $1 }')
# 删除docker
docker rm $(docker ps -a | grep "Exited" | awk '{print $1 }')
# 删除images
docker rmi $(docker images | grep "none" | awk '{print $3}')