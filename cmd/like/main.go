package main

import (
	service "dousheng/cmd/like/internal/service"
	like "dousheng/kitex_gen/like/likeservice"
	"dousheng/pkg/database"
	"dousheng/pkg/etcd_discovery"
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
	tracer.InitJaeger(g.ServiceLikeName)
	etcd_discovery.InitUserRpc()
	etcd_discovery.InitVideoRpc()
	database.InitDB()
	service.CronTaskSetUp() // 定时刷新
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", g.ServiceLikeAddress) // TODO 写自己的服务地址
	r, err := etcd.NewEtcdRegistry([]string{g.EtcdAddress})    // r should not be reused.
	if err != nil {
		log.Println(err.Error())
	}
	svr := like.NewServer(new(LikeServiceImpl), //TODO 新建自己的服务
		server.WithServiceAddr(addr),                          // 定义端口
		server.WithSuite(opentracing.NewDefaultServerSuite()), // 链路监听
		server.WithMuxTransport(),                             // 多路复用
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: g.ServiceLikeName}), // TODO 写自己的服务名
		server.WithRegistry(r), // 注册服务
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
