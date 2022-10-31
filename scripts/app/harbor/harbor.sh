#!/bin/bash
#
#***********************************************************************
#         Author:            1ch0
#         Date:              2022-01-07
#         FileName:          harbor_install.sh
#         Description:       1ch0 script
#         Blog:               https://1ch0.github.io/
#         Copyright (C): 2022 All rights reserved
#***********************************************************************
apt install openssl

curl -LO https://hub.fastgit.org/goharbor/harbor/releases/download/v2.2.2/harbor-online-installer-v2.2.2.tgz
tar -zxvf harbor-online-installer-v2.2.2.tgz
cd harbor && ls -la

## 自建 CA
openssl genrsa -out ca.key 4096

openssl req -x509 -new -nodes -sha512 -days 3650 \
    -subj "/C=CN/ST=Shenzhen/L=Shenzhen/O=example/OU=Personal/CN=k8scat.com" \
    -key ca.key \
    -out ca.crt
## 生成域名证书
openssl genrsa -out harbor.k8scat.com.key 4096
## 生成证书签名请求文件 CSR（Certificate Signing Request）
openssl req -sha512 -new \
    -subj "/C=CN/ST=Shenzhen/L=Shenzhen/O=example/OU=Personal/CN=harbor.k8scat.com" \
    -key harbor.k8scat.com.key \
    -out harbor.k8scat.com.csr
## 生成 x509 v3 扩展文件，以此来满足 SAN（Subject Alternative Name） 和 x509 v3 扩展的要求
cat > v3.ext <<-EOF
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
extendedKeyUsage = serverAuth
subjectAltName = @alt_names

[alt_names]
DNS.1=harbor.k8scat.com
EOF
## 使用 ca.crt、ca.key、harbor.k8scat.com.csr 和 v3.ext 来生成我们需要的域名证书
openssl x509 -req -sha512 -days 3650 \
    -extfile v3.ext \
    -CA ca.crt -CAkey ca.key -CAcreateserial \
    -in harbor.k8scat.com.csr \
    -out harbor.k8scat.com.crt
## 配置 Harbor 和 Docker 的证书
mkdir -p /data/cert/
cp harbor.k8scat.com.crt /data/cert/
cp harbor.k8scat.com.key /data/cert/

## Docker
## 转换
openssl x509 -inform PEM -in harbor.k8scat.com.crt -out harbor.k8scat.com.cert
mkdir -p /etc/docker/certs.d/harbor.k8scat.com/
cp harbor.k8scat.com.cert /etc/docker/certs.d/harbor.k8scat.com/
cp harbor.k8scat.com.key /etc/docker/certs.d/harbor.k8scat.com/
cp ca.crt /etc/docker/certs.d/harbor.k8scat.com/
systemctl restart docker

