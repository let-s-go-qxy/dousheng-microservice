
### 开发流程
1. 在src下新建微服务文件夹，进入后打开终端执行：
   kitex -module dousheng -I ../../idl/ -use dousheng/kitex_gen  -service xxx ../../idl/xx.proto
2. 完善main函数的**服务注册**，**init函数**，和**启动服务器**代码，完善handler服务的业务
3. init 函数中包含初始化数据库，初始化链路追踪，按需求是否要初始化服务发现
4. 如果需要调用到其他服务，新建rpc目录实现服务发现

### 业务开发后生产执行流程
4. 在该目录下开发业务，业务完成后可以使用下面两步生成可执行文件：
5. sh build.sh   这里不出错会出现 ouput 文件夹
6. sh output/bootstrap.sh 

### win 启动方法
1. 启动etcd：
   [使用二进制安装etcd](https://blog.csdn.net/wohu1104/article/details/115794367)，
   使用 `/tmp/etcd-download-test/etcd` 命令
2. 启动各项目

// 用户模块
kitex -module dousheng -I ../../idl/ -use dousheng/kitex_gen  -service userservice ../../idl/user.proto

// 社交模块
kitex -module dousheng -I ../../idl/ -use dousheng/kitex_gen  -service relationservice ../../idl/relation.proto


