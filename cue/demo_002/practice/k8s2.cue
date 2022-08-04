// 导入k8s资源模块，并创建资源

package kube
import (
  "k8s.io/api/core/v1"
  "k8s.io/api/extensions/v1beta1"
)
service <Name>: v1.Service
deployment <Name>: v1beta1.Deployment