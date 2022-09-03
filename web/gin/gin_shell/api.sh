#!/bin/bash

kill -9 $(netstat -nlp | grep :8080 | awk '{print $7}' | awk -F"/" '{ print $1 }')
cd /root/workspace/app/demo/go/web/gin/reponse
nohup go run main.go &

curl -H "Content-Type:application/json" -X POST -d '{"token":"xReadGroupArgs"}' http://124.223.36.219:8888/api
