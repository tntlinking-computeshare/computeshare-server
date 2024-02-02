# 环境安装文档

## 基础环境安装

### kubernetes

#### kubernetes 安装
[kubernetes傻瓜式安装](https://github.com/abing258/kubernetes-study/blob/main/Linux-master-node-installation.md)

### 基础环境安装
```shell
kubectl apply -f base-service.yml

##启动成功后创建数据库
create DATABASE computeshare;

## 初始化数据 init.sql
## 运行项目中的 TestAddGatewaysPorts 生成端口数据
```

### seaweedfs 中心
1. 主节点的运行
```shell
##使用项目中的docker-compose.yml 修改ip
docker compose up -d 

```

## computeshare server 服务安装
```shell
## 按需要修改端口设置和域名配置
kubectl apply -f computeshare-server.yml
```

## computeshare client 安装


### computeshare client 客户端

### frp 端口穿透服务安装

```shell
wget https://github.com/fatedier/frp/releases/download/v0.53.2/frp_0.53.2_linux_amd64.tar.gz
tar -xzvf frp_0.53.2_linux_amd64.tar.gz
mv frp_0.53.2_linux_amd64/frps /usr/local/bin/

mkdir -p /etc/frp/

mv frp_0.53.2_linux_amd64/frps.toml /etc/frp/

cat > /lib/systemd/system/frps.service <<EOF
[Unit]

Description=frps

[Service]

PIDFile=/run/frps.pid

#EnvironmentFile=/etc/systemd/test.conf

ExecStart=/usr/local/bin/frps -c /etc/frp/frps.toml 

ExecReload=/bin/kill -SIGHUP $MAINPID

ExecStop=/bin/kill -SIGINT $MAINPID


[Install]

WantedBy=multi-user.target
EOF

systemctl daemon-reload

systemctl enable frps
systemctl start frps

```

将frps 服务7000端口映射到公网， computehsare-client 客户端 server.p2p 配置 frps的公网和端口
