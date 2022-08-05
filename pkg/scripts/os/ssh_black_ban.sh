#!/bin/bash
#
#***********************************************************************
#         Author:            1ch0
#         Date:              2022-01-27
#         FileName:          ssh_black_ban.sh
#         Description:       1ch0 script
#         Blog:              https://1ch0.github.io/
#         Copyright (C):     2022 All rights reserved
#***********************************************************************
#set -x
#
DEFINE="10"
for i in `cat b.txt`; do
  echo $i
  #NUM="`echo $i|awk '{print $1}'`"
  #echo $NUM
  #IP="`echo $i|tr -s ' '|cut -d ' ' -f3`"
  #if [$NUM -gt $DEFINE]
 # then
  #  EXIST="`grep $IP /etc/hosts.deny|wc -l`"
   # if [$EXIST -gt 0]
    #then
    #  echo "sshd: $IP" >> /etc/hosts.deny
   # fi
 # fi
done
