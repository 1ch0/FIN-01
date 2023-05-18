#!/usr/bin/env bash

helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update bitnami
helm search repo mongodb

helm pull bitnami/mongodb --version 13.6.0
tar -xvf mongodb-13.6.0
