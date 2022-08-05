#!/bin/bash
#
#***********************************************************************
#         Author:            1ch0
#         Date:              2022-01-27
#         FileName:          sec_1ch0.sh
#         Description:       1ch0 script
#         Blog:              https://1ch0.github.io/
#         Copyright (C):     2022 All rights reserved
#***********************************************************************
#set -x
#
echo "主机安全检测脚本"

echo "警告：本脚本只是一个检查的操作，未对服务器做任何修改，管理员可以根据此报告进行相应的设置。"

echo ---------------------------------------主机安全检查-----------------------

echo "系统版本"

uname -a

echo --------------------------------------------------------------------------

echo "本机的ip地址是："

ifconfig | grep --color "\([0-9]\{1,3\}\.\)\{3\}[0-9]\{1,3\}"

echo --------------------------------------------------------------------------

awk -F":" '{if($2!~/^!|^*/){print "("$1")" " 是一个未被锁定的账户，请管理员检查是否需要锁定它或者删除它。"}}' /etc/shadow

echo --------------------------------------------------------------------------

more /etc/login.defs | grep -E "PASS_MAX_DAYS" | grep -v "#" |awk -F' '  '{if($2!=90){print "/etc/login.defs里面的"$1 "设置的是"$2"天，请管理员改成90天。"}}'

echo --------------------------------------------------------------------------

more /etc/login.defs | grep -E "PASS_MIN_LEN" | grep -v "#" |awk -F' '  '{if($2!=6){print "/etc/login.defs里面的"$1 "设置的是"$2"个字符，请管理员改成6个字符。"}}'

echo -------------------------------------------------------------------------

more /etc/login.defs | grep -E "PASS_WARN_AGE" | grep -v "#" |awk -F' '  '{if($2!=10){print "/etc/login.defs里面的"$1 "设置的是"$2"天，请管理员将口令到期警告天数改成10天。"}}'

echo --------------------------------------------------------------------------

grep TMOUT /etc/profile /etc/bashrc > /dev/null|| echo "未设置登录超时限制，请设置之，设置方法：在/etc/profile或者/etc/bashrc里面添加TMOUT=600参数"

echo --------------------------------------------------------------------------

if ps -elf |grep xinet |grep -v "grep xinet";then

echo "xinetd 服务正在运行，请检查是否可以把xinnetd服务关闭"

else

echo "xinetd 服务未开启"

fi

echo --------------------------------------------------------------------------

echo "查看系统密码文件修改时间"

ls -ltr /etc/passwd

echo --------------------------------------------------------------------------

echo  "查看是否开启了ssh服务"

if service sshd status | grep -E "listening on|active \(running\)"; then

echo "SSH服务已开启"

else

echo "SSH服务未开启"

fi

echo --------------------------------------------------------------------------

echo "查看是否开启了TELNET服务"

if more /etc/xinetd.d/telnetd 2>&1|grep -E "disable=no"; then

echo  "TELNET服务已开启 "

else

echo  "TELNET服务未开启 "

fi

echo --------------------------------------------------------------------------

echo  "查看系统SSH远程访问设置策略(host.deny拒绝列表)"

if more /etc/hosts.deny | grep -E "sshd: ";more /etc/hosts.deny | grep -E "sshd"; then

echo  "远程访问策略已设置 "

else

echo  "远程访问策略未设置 "

fi

echo --------------------------------------------------------------------------
