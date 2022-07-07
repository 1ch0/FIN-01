#!/bin/bash
set +x
echo '---------------------------------------------------'
ps -eo user,pid,pcpu,pmem,args --sort=-pcpu  |head -n 10
echo '---------------------------------------------------'
ps -eo user,pid,pcpu,pmem,args --sort=-pmem  |head -n 10
