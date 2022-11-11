#!/usr/bin/env bash
wget https://studygolang.com/dl/golang/go1.19.3.linux-amd64.tar.gz

rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.3.linux-amd64.tar.gz

tee -a $HOME/.bashrc <<'EOF'
# Go envs
export PATH=$PATH:/usr/local/go/bin
export GO111MODULE="on" # 开启 Go moudles 特性
export GOPROXY=https://goproxy.cn,direct # 安装 Go 模块时，代理服务器设置
export G_MIRROR=https://golang.google.cn/dl/
export GOSUMDB=off # 关闭校验 Go 依赖包的哈希值
EOF
source ~/.bashrc
go version