package etcd_discovery

import (
	"dousheng/kitex_gen/user/userservice"
	g "dousheng/pkg/global"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	tracing "github.com/kitex-contrib/tracer-opentracing"
)

var UserClient userservice.Client

// InitUserRpc 服务发现
func InitUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{g.EtcdAddress})
	if err != nil {
		panic(err)
	}
	cli, err := userservice.NewClient(
		g.ServiceUserName,                                 // TODO 这里需要改成自己的服务名
		client.WithSuite(tracing.NewDefaultClientSuite()), // tracer
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	UserClient = cli
}
