#!/bin/bash
#
#***********************************************************************
#         Author:            1ch0
#         Date:              2022-01-06
#         FileName:          cloudreve_install.sh
#         Description:       1ch0 script
#         Blog:               https://1ch0.github.io/
#         Copyright (C): 2022 All rights reserved
#***********************************************************************

wget -c https://hub.fastgit.org/cloudreve/Cloudreve/releases/download/3.4.2/cloudreve_3.4.2_linux_amd64.tar.gz
# 解压程序包
tar -zxvf cloudreve_3.4.2_linux_amd64.tar.gz 
rm -f cloudreve_3.4.2_linux_amd64.tar.gz
# 赋予执行权限
chmod +x cloudreve

mv cloudreve /usr/bin/
# 启动 Cloudreve
cloudreve
