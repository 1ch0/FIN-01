# [kind](https://kind.sigs.k8s.io/)

### 创建上述集群：
kind create cluster --name multi-node --config=kind-config.yaml
### 切换到该集群：
kubectl cluster-info --context kind-multi-node