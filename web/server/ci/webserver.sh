#!/bin/bash

project_path="git@webserver.git"
deploy_path="/root/webserver"
git clone $project_path $deploy_path
cd $deploy_path

kill -9 $(netstat -nlp | grep :8080 | awk '{print $7}' | awk -F"/" '{ print $1 }')
git fetch --all
git reset --hard origin/develop
git pull
git checkout develop

go mod tidy
nohup go run main.go &
