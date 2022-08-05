#!/usr/bin/env bash
wget https://studygolang.com/dl/golang/go1.17.11.linux-amd64.tar.gz -O /tmp/go1.17.11.linux-amd64.tar.gz

mkdir -p $HOME/go
mkdir -p $HOME/gopath
tar -xvzf /tmp/go1.17.11.linux-amd64.tar.gz -C $HOME/go
mv $HOME/go/go $HOME/go/go1.17.11
rm -f /tmp/go1.17.11.linux-amd64.tar.gz

tee -a $HOME/.bashrc <<'EOF'
# Go envs
export GOVERSION=go1.17.11 # Go 版本设置
export GO_INSTALL_DIR=$HOME/go # Go 安装目录
export GOROOT=$GO_INSTALL_DIR/$GOVERSION # GOROOT 设置
export GOPATH=$HOME/gopath # GOPATH 设置
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH # 将 Go 语言自带的和通过 go install 安装的二进制文件加入到 PATH 路径中
export GO111MODULE="on" # 开启 Go moudles 特性
export GOPROXY=https://goproxy.cn,direct # 安装 Go 模块时，代理服务器设置
export G_MIRROR=https://golang.google.cn/dl/
export GOPRIVATE=
export GOSUMDB=off # 关闭校验 Go 依赖包的哈希值
EOF
source ~/.bashrc
go version