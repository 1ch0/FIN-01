#!/bin/bash
#
#***********************************************************************
#         Author:            1ch0
#         Date:              2022-01-27
#         FileName:          1ch0_sec.sh
#         Description:       1ch0 script
#         Blog:              https://1ch0.github.io/
#         Copyright (C):     2022 All rights reserved
#***********************************************************************
set -x
#
echo --------------------------------主机安全检查-----------------------
echo "警告：本脚本只是一个检查的操作，未对服务器做任何修改，管理员可以根据此报告进行相应的设置。"



echo -e "\033[41;34m 系统版本\033[0m"
uname -a

echo -e "\033[41;34m 本机的ip地址\033[0m"
ifconfig | grep --color "\([0-9]\{1,3\}\.\)\{3\}[0-9]\{1,3\}"

echo -e "\033[41;34m 身份鉴别\033[0m"

echo -e "\033[42;34m cat /etc/shadow \033[0m"
cat /etc/shadow

echo -e "\033[42;34m cat /etc/login.defs|grep MAX_DAY|grep -v \# \033[0m"
cat /etc/login.defs|grep MAX_DAY|grep -v \# 

echo -e "\033[42;34m 密码复杂度 \# \033[0m"
cat /etc/pam.d/system-auth|grep quality

echo -e "\033[41;34m 访问控制\033[0m"
awk -F":" '{if($2!~/^!|^*/){print "("$1")" " 是一个未被锁定的账户，请管理员检查是否需要锁定它或者删除它。"}}' /etc/shadow

echo -e "\033[42;34m cat /etc/passwd\033[0m"
cat /etc/passwd

echo -e "\033[42;34m cat /etc/shadow\033[0m"
cat /etc/shadow

echo -e "\033[42;34m id root\033[0m"
id root

echo -e "\033[42;34m ls -l /etc/|grep passwd|grep shadow|grep profile\033[0m"
ls -l /etc/passwd
ls -l /etc/shadow
ls -l /etc/profile
