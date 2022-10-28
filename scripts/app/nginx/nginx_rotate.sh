#!/bin/bash
#
#***********************************************************************
#         Author:            1ch0
#         Date:              2022-08-26
#         FileName:          rotate.sh
#         Description:       1ch0 script
#         Blog:              https://1ch0.github.io/
#         Copyright (C):     2022 All rights reserved
#***********************************************************************
set -x
#
LOGS_PATH=/usr/local/nginx/logs/history
CUR_LOGS_PATH=/usr/local/nginx/logs
YESTERDAY=$(date -d "yesterday" +%Y-%m-%d)
mv ${CUR_LOGS_PATH}/dev01_access.log ${LOGS_PATH}/dev01_access_${YESTERDAY}.log
mv ${CUR_LOGS_PATH}/dev02_access.log ${LOGS_PATH}/dev02_access_${YESTERDAY}.log
mv ${CUR_LOGS_PATH}/error.log ${LOGS_PATH}/error_${YESTERDAY}.log
## 向 Nginx 主进程发送 USR1 信号。 USR1 信号是重新打开日志文件
kill -USR1 $(cat /usr/local/nginx/logs/nginx.pid)