#!/bin/bash

diskUsage=`df | grep '/$'| awk '{print $(NF-1)}' | awk -F'%' '{print $1}'`
if [ $diskUsage -ge 80 ] ;then
docker rmi -f $(docker images | grep "none" | awk '{print $3}')
fi