#!/bin/bash

project_path="git@apiserver.git"
deploy_path="/root/apiserver"
git clone $project_path $deploy_path
cd $deploy_path

kill -9 $(netstat -nlp | grep :8000 | awk '{print $7}' | awk -F"/" '{ print $1 }')
git fetch --all
git reset --hard origin/develop
git pull
git checkout develop

nohup make run &
