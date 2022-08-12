#!/bin/bash
#
#***********************************************************************
#         Author:            1ch0
#         Date:              2022-01-07
#         FileName:          scp-key.sh
#         Description:       1ch0 script
#         Blog:               https://1ch0.github.io/
#         Copyright (C): 2022 All rights reserved
#***********************************************************************
IP="
172.31.7.101
172.31.7.102
172.31.7.103
172.31.7.104
172.31.7.105
172.31.7.106
"

for node in ${IP};do
  sshpass -p 123456 ssh-copy-id ${node} -o StrictHostKeyChecking=no
  if [ $? -eq 0 ];then
    echo "${node} 密钥 copy 完成"
  else
    echo "${node} 密钥 copy 失败"
  fi
done
