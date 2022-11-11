#!/bin/bash
set -x
lsb_release -a
mkdir $HOME/workspace
mkdir $HOME/workspace/github.com
mkdir $HOME/workspace/app
mkdir $HOME/workspace/app/docker-compose
mkdir $HOME/workspace/app/yaml
mkdir $HOME/workspace/app/scripts


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




cp /etc/apt/sources.list /etc/apt/sources.list_backup

tee /etc/apt/sources.list << EOF
#阿里源
deb http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse
EOF
apt-get -y update
apt-get -y upgrade
sudo apt-get install libc6-dev

# 添加公钥
apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 40976EAF437D05B5

apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 3B4FE6ACC0B21F32

apt-get -y update
apt-get -y upgrade
apt -y autoremove

apt install -y screenfetch git wget net-tools vim libc6-dev snapd snapcraft add-apt-repository xclip



# 安装 neovim
sudo add-apt-repository ppa:neovim-ppa/unstable
sudo apt-get update
sudo apt-get install -y neovim
# 安装 spacevim
# curl -sLf https://spacevim.org/cn/install.sh | bash

# 使mkfontscale和mkfontdir命令正常运行
sudo apt-get install -y ttf-mscorefonts-installer
# 使fc-cache命令正常运行
sudo apt-get install -y fontconfig

tee -a $HOME/.vim/vimrc << EOF
autocmd BufNewFile *.sh exec ":call SetTitle()"
func SetTitle()
    if expand("%:e") == 'sh'
    call setline(1, "#!/bin/bash")
    call setline(2, "set -x")
    call setline(3, "")
    call setline(4, "#Author:       1ch0")
    call setline(5, "#Date:         ".strftime("%Y-%m-%d"))
    call setline(6, "#FileName:     ".expand("%"))
    call setline(7, "#Description:  1ch0 script")
    call setline(8, "#Blog:         https://1ch0.github.io/")
    call setline(9, "#Copyright(C): ".strftime("%Y")." All rights reserved")
    call setline(10, "")
    call setline(11, "")
    endif
endfunc
autocmd BufNewFile * normal G
EOF
