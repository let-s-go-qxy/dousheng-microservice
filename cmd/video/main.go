package main

import (
	video "dousheng/kitex_gen/video/videoservice"
	"dousheng/pkg/database"
	"dousheng/pkg/etcd_discovery"
	g "dousheng/pkg/global"
	"dousheng/pkg/oss_init"
	"dousheng/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"net"
)

func init() {
	tracer.InitJaeger(g.ServiceVideoName)
	initRpc()
	database.InitSpecificDB()
	database.InitVideoRedis()
	oss_init.OSSInit()
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", g.ServiceVideoAddress) // TODO 写自己的服务地址
	r, err := etcd.NewEtcdRegistry([]string{g.EtcdAddress})     // r should not be reused.
	if err != nil {
		log.Println(err.Error())
	}
	svr := video.NewServer(new(VideoServiceImpl),
		server.WithServiceAddr(addr),                          // 定义端口
		server.WithSuite(opentracing.NewDefaultServerSuite()), // 链路监听
		server.WithMuxTransport(),                             // 多路复用
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: g.ServiceVideoName}), // TODO 写自己的服务名
		server.WithRegistry(r), // 注册服务
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}

func initRpc() {
	etcd_discovery.InitUserRpc()
	etcd_discovery.InitVideoRpc()
	etcd_discovery.InitRelationRpc()
	etcd_discovery.InitMessageRpc()
	etcd_discovery.InitLikeRpc()
	etcd_discovery.InitCommentRpc()
}
