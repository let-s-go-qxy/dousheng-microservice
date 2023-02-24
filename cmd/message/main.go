package main

import (
	"dousheng/cmd/message/internal/service"
	"dousheng/cmd/message/num"
	message "dousheng/kitex_gen/message/messageservice"
	"dousheng/pkg/database"
	g "dousheng/pkg/global"
	"dousheng/pkg/mq"
	"dousheng/pkg/tracer"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"net"
	"time"
)

var (
	producer1 rocketmq.Producer
	consumer1 rocketmq.PushConsumer
)

func init() {
	tracer.InitJaeger(g.ServiceMessageName)
	database.InitDB()
	rlog.SetLogLevel("error")
	var err error
	producer1, consumer1, err = mq.InitRM()
	if err != nil {
		panic(err.Error())
	}
	// 初始化MQ
	mq.SubscribeMessage(service.SubscribeMessageAction)
}

func main() {
	// 关闭队列
	defer func(producer1 rocketmq.Producer) {
		err := producer1.Shutdown()
		err = consumer1.Shutdown()
		if err != nil {
			hlog.Error(err.Error())
		}
	}(producer1)
	addr, _ := net.ResolveTCPAddr("tcp", g.ServiceMessageAddress) // TODO 写自己的服务地址
	r, err := etcd.NewEtcdRegistry([]string{g.EtcdAddress})       // r should not be reused.
	if err != nil {
		log.Println(err.Error())
	}
	svr := message.NewServer(new(MessageServiceImpl),
		server.WithServiceAddr(addr),                          // 定义端口
		server.WithSuite(opentracing.NewDefaultServerSuite()), // 链路监听
		server.WithMuxTransport(),                             // 多路复用
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: g.ServiceMessageName}), // TODO 写自己的服务名
		server.WithRegistry(r), // 注册服务
	)
	go func() {
		for true {
			fmt.Println("orm num:", num.ORMNum, " message num", num.MessageNum, "req num", num.ReqNUM)
			time.Sleep(time.Second * 3)
		}
	}()
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
