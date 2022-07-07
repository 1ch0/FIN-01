#!/bin/bash
mkdir $HOME/workspace

tee -a $HOME/.bashrc << EOF
# User specific environment
# Basic envs
alias rm='rm -i'
alias cp='cp -i'
alias mv='mv -i'
alias ll='ls -al'
alias ch='chmod a+x'
alias cip='curl cip.cc'
alias aip='curl -s http://myip.ipip.net'
alias c='clear'
alias vi='nvim'
# alias psi='reptyr -s'
alias psg='ps -aux|grep'
alias s='screenfetch'
alias x='xclip'
alias k='kubectl'
alias ks='kubectl -n kube-system'
alias ka='kubectl apply -f'
alias kd='kubectl delete -f'
alias dc='docker-compose'


export LANG="en_US.UTF-8" # 设置系统语言为 en_US.UTF-8，避免终端出现中文乱码
export PS1='[\u@dev \W]\$ ' # 默认的 PS1 设置会展示全部的路径，为了防止过长，这里只展示："用户名@dev 最后的目录名"


export WORKSPACE="$HOME/workspace" # 设置工作目录
export PATH=$HOME/bin:$PATH # 将 $HOME/bin 目录加入到 PATH 变量中
# Default entry folder
cd \$WORKSPACE
# 登录系统，默认进入 workspace 目录
EOF

source  $HOME/.bashrc