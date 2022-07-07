#!/bin/bash
#
#***********************************************************************
#         Author:            1ch0
#         Date:              2022-01-06
#         FileName:          docker_install.sh
#         Description:       1ch0 script
#         Blog:               https://1ch0.github.io/
#         Copyright (C): 2022 All rights reserved
#***********************************************************************
# 卸载旧版
sudo apt-get remove docker docker-engine docker.io containerd runc
# 安装 docker 1
curl -fsSL https://get.docker.com | bash -s docker --mirror aliyun
# 安装 docker 2
curl -sSL https://get.daocloud.io/docker | sh

# 安装 docker compose
curl -L https://get.daocloud.io/docker/compose/releases/download/1.25.4/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
docker-compose -v
systemctl daemon-reload
service docker restart
service docker status 
