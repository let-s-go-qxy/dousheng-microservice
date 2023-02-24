
### 开发流程
1. 在src下新建微服务文件夹，进入后打开终端执行：
   kitex -module dousheng -I ../../idl/ -use dousheng/kitex_gen  -service xxx ../../idl/xx.proto
2. 完善main函数的**init函数**，**服务注册**，和**启动服务器**代码，完善handler服务的业务
3. init 函数中包含初始化数据库，初始化链路追踪，按需求是否要初始化服务发现

// 用户模块
kitex -module dousheng -I ../../idl/ -use dousheng/kitex_gen  -service userservice ../../idl/user.proto

// 社交模块
kitex -module dousheng -I ../../idl/ -use dousheng/kitex_gen  -service relationservice ../../idl/relation.proto

### 业务开发后生产执行流程
4. 在该目录下开发业务，业务完成后可以使用下面两步生成可执行文件：
5. sh build.sh   这里不出错会出现 ouput 文件夹
6. sh output/bootstrap.sh 

### 火焰图
go tool pprof -http :8080 localhost:9000/debug/pprof/profile?seconds=30


