package main

import (
	relation "dousheng/kitex_gen/relation/relationservice"
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
	tracer.InitJaeger(g.ServiceRelationName)
	etcd_discovery.InitUserRpc()
	etcd_discovery.InitMessageRpc()
	database.InitDB()
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", g.ServiceRelationAddress) // TODO 声明自己的服务地址
	r, err := etcd.NewEtcdRegistry([]string{g.EtcdAddress})        // r should not be reused.
	if err != nil {
		log.Println(err.Error())
	}
	svr := relation.NewServer(new(RelationServiceImpl),
		server.WithServiceAddr(addr),                          // 定义端口
		server.WithMuxTransport(),                             // 多路复用
		server.WithSuite(opentracing.NewDefaultServerSuite()), // 链路监听
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: g.ServiceRelationName}), // TODO 声明自己的服务名
		server.WithRegistry(r), // 注册服务 -> etcd
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
