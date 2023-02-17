package main

import (
	comment "dousheng/kitex_gen/comment/commentservice"
	g "dousheng/pkg/global"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", g.ServiceCommentAddress) // TODO 写自己的服务地址
	r, err := etcd.NewEtcdRegistry([]string{g.EtcdAddress})       // r should not be reused.
	if err != nil {
		log.Println(err.Error())
	}
	svr := comment.NewServer(new(CommentServiceImpl), //TODO 新建自己的服务
		server.WithServiceAddr(addr),                          // 定义端口
		server.WithSuite(opentracing.NewDefaultServerSuite()), // 链路监听
		server.WithMuxTransport(),                             // 多路复用
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: g.ServiceCommentName}), // TODO 写自己的服务名
		server.WithRegistry(r), // 注册服务
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
