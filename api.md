# 共享算力API接口

```
接口地址：https://computeshare.newtouch.com/api
授权认证：
  使用http header的标准Authorization认证
  Authorization: Bearer {{ access_token }}
```

另外也可以参考在线swagger文档： https://api.computeshare.newtouch.com/q/swagger-ui

## 认证

### 获取认证Token

此接口不需要认证， 此处的认证接口获取的token有效期是10年
#### Request
```shell
### 获取认证token
POST https://computeshare.newtouch.com/api/v1/user/loginWithClient
Content-Type: application/json

{
  "username": "1300xxxxxxx",
  "password": "somePassword"
}

```

参数说明：
* username: 用户名，默认是手机号码，可以在页面上查看和修改
* password: 密码

#### Response
成功响应：
```
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "someToken"
  }
}
```

参数说明：
* token: 认证后的token，用于后续接口的认证

## 计算资源

### 启动沙箱环境

一次性创建一台虚拟机，并为其配置网络映射


#### Request

```shell
### 沙盒流程测试
POST https://api.computeshare.newtouch.com/v1/sandbox
Content-Type: application/json
Authorization: Bearer {{ access_token }}

{
  "instance": {
    "specId": 1,
    "imageId": 1,
    "duration": 1,
    "name": "mohaijiang-process1",
    "publicKey": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC2mLWYddGeahdk6i3muy72XDbppnG4LIDhyj/rSuzLstdVLI7mF7efkwCZgyYcYRJoIjNI5mnb17o7/qVWdgGSiMnSgiPcw4r0Dp1pghWXBEog3o7pI3gicY6//Y4+liqypBEDmBSJnDsMJqVARzFV0rjJLhYSCbYk99LPB1ZLj0mDvIY/1SjRR9bfPuW9Ht6QjkS9DEWIdTrJ0dAaGwJkc+a5pCVzcopq4ycvBVLEnEq4xCrhbNx/LrpYxytA7WXg6kUcN+4Me63QVPxUExcn14qXr5uYxo+ePkoBCNdbqFsm0Z1rxrEX8oGDHvAfsoCpQr/OV8J5WwO7i/QIOyK7 mohaijiang110@163.com",
    "password": "Abcd1234",
    "dockerCompose": "c2VydmljZXM6CiAgcHJveHk6CiAgICBpbWFnZTogbmdpbngKICAgIHBvcnRzOgogICAgICAtIDgwOjgw"
  },
  "networkMapping": [
    {
      "name": "nginx",
      "protocol": "TCP",
      "computerPort": 80
    }
  ]
}
```


参数说明：
* instance: 资源实例配置
  * specId: 资源规格，1：2c4g, 2:4c8g
  * imageId: 镜像id: 1: ubuntu:20.04
  * duration: 资源有效期(天)
  * name: 资源名
  * publicKey: 用户登录公钥
  * password: 用户密码
  * dockerCompose： base64格式化后的docker-compose.yaml, 启动虚拟机后，会自动启动这个docker-compose.yml 文件
* networkMapping: 网络映射数组
  * name: 网络映射名
  * protocol: 映射协议,当前仅支持TCP
  * computerPort: 内网映射出的端口
#### Response
成功响应：
```
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "someToken"
  }
}
