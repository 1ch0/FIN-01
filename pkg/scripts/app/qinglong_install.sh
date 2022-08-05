#!/bin/bash
#
#***********************************************************************
#         Author:            1ch0
#         Date:              2022-01-19
#         FileName:          qinglong_install.sh
#         Description:       1ch0 script
#         Blog:               https://1ch0.github.io/
#         Copyright (C): 2022 All rights reserved
#***********************************************************************
docker run -dit \
  -v /data/QL/config:/ql/config \
  -v /data/QL/log:/ql/log \
  -v /data/QL/db:/ql/db \
  -v /data/QL/repo:/ql/repo \
  -v /data/QL/raw:/ql/raw \
  -v /data/QL/scripts:/ql/scripts \
  -v /data/QL/jbot:/ql/jbot \
  -v /data/QL/deps:/ql/deps \
  -p 59700:5700 \
  --name QL \
  --hostname QL \
  --restart unless-stopped \
  whyour/qinglong:latest
