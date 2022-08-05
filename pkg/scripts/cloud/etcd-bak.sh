#!/bin/bash
#
#***********************************************************************
#         Author:            1ch0
#         Date:              2022-02-17
#         FileName:          etcd-bak.sh
#         Description:       1ch0 script
#         Blog:              https://1ch0.github.io/
#         Copyright (C):     2022 All rights reserved
#***********************************************************************
set -x
#
# TODO mkdir -p /data/etcd-backup-file/
source /etc/profile
DATE=`date+%Y-%m-%d_%H-%M-%S`
ETCDCTL_API=3 /usr/bin/etcdctl snapshot save /data/etcd-backup-file/etcd-snapshot-${DATE}.db
