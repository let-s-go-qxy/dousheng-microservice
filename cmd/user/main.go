package main

import (
	user "dousheng/kitex_gen/user/userservice"
	"dousheng/pkg/database"
	g "dousheng/pkg/global"
	"dousheng/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"net"
)

func init() {
	tracer.InitJaeger(g.ServiceUserName)
	database.InitDB()
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", g.ServiceUserAddress)
	r, err := etcd.NewEtcdRegistry([]string{g.EtcdAddress}) // r should not be reused.
	if err != nil {
		log.Println(err.Error())
	}
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),                          // 定义端口
		server.WithSuite(opentracing.NewDefaultServerSuite()), // 链路监听
		server.WithMuxTransport(),                             // 多路复用
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: g.ServiceUserName}), // 声明服务名
		server.WithRegistry(r), // 注册服务
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
