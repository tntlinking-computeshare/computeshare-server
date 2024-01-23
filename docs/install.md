# 环境安装文档

## 基础环境安装

### kubernetes 

#### kubernetes 安装

#### helm 安装

#### nginx ingress 安装

```shell
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install ingress-nginx ingress-nginx/ingress-nginx --set controller.hostPort.enabled=true --namespace ingress-nginx --create-namespace
```

#### cert-manager 安装

```shell
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm install   cert-manager jetstack/cert-manager   --namespace cert-manager   --create-namespace   --version v1.13.2   --set installCRDs=true
```

#### longhorn 安装

```shell
helm repo add longhorn https://charts.longhorn.io
helm repo update
helm install longhorn longhorn/longhorn --namespace longhorn-system --create-namespace --version 1.5.3
```

### seaweedfs 中心

## computeshare server 服务安装


## computeshare client 安装

### 物理机基础环境

### computeshare client 客户端

### 
