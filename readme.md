
### 目录结构
``` shell
dousheng-microservice
├─ cmd # 业务开发代码目录
├─ conf # 配置目录，不应该缓存到git
├─ idl # idl目录
├─ kitex_gen # kitex的代码生成目录
└─ pkg # 各种全局可用的包，依赖，工具
    ├─ global # 常量，错误码存放地 （为了配合之前写的，改名为global）
    │   ├─ errno # 错误码存放地
    │   ├─ define # 常量定义
    │   └─ rpc # 微服务相关
    ├─ utils # 自定义工具包
    ├─ database # 数据库相关
    ├─ etcd_discovery # etcd注册
    └─ tracer # 链路追踪相关
```

### 开发流程
1. 编写 idl
2. 生成 idl
   kitex -module dousheng -I ./idl/ ./idl/xx.proto
3. 完善各类配置（如 pkg/global/rpc.go ）
4. 在 cmd 目录下生成对应服务...（详情见src下readme）
5. 再 pkg/etcd_discovery/ 目录下新增对应微服务的发现中心（只需要某个服务发现的服务可以写在自己的微服务下）
6. 利用 docker-compose 启动 etcd，jaeger : `sudo docker-compose up`
7. 服务完善后，依次启动微服务，最后启动 api_gateway 网关服务
8. 使用模拟器或 apifox 测试接口， 使用 http://127.0.0.1:16686/search 查看链路追踪
   ![链路追踪预览图](https://aeiblog-1301396258.cos.ap-chengdu.myqcloud.com/img/20230216222404.png)

### 环境准备
#### docker
win用户在本地的docker desktop中配置：
打开docker desktop，进入Settings > Resources > WSL Integration，激活Enable integration with additional distros:Ubuntu-20.04，Apply&Restart。
#### docker-compose
[dokcer-compose 安装](https://cloud.tencent.com/developer/article/2204414)

使用 docker 配置 etcd 和 jaeger :
```shell
sudo docker-compose up # 项目根目录终端运行
```

### tips:
- 使用proto import功能时，要在settings -> Protocol Buffers 增加idl目录
- docker使用过程中可能会遇到的问题和解决方案：[WSL DETECTED: We recommend using Docker Desktop for Windows](https://blog.csdn.net/yuezhilanyi/article/details/117036433)

### 微服务交互步骤
1. handleer.go 实现我们的方法
2. main.go 启动微服务，注册到etcd
3. etcd_discovery 生成服务发现客户端，如果只开放部分方法，不公开 client，只公开方法。
4. 另一个服务在自己的 main.go init函数中完成服务发现，并且调用客户端或者方法。