#!/bin/bash
curl -fsSl https://static.kubevela.net/script/install-velad.sh | bash -s 1.4.3
velad uninstall
velad install
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml
vela comp
# vela addon enable ~/.vela/addons/velaux
# use db
vela addon enable ~/.vela/addons/velaux dbType=mongodb database=kubevela dbURL=mongodb://root:pwd@10.33.32.67:27017
nohup kubectl port-forward --address 0.0.0.0 service/velaux -n vela-system 58080:80 > /dev/null 2>&1 &
exit