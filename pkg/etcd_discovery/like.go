package etcd_discovery

import (
	"dousheng/kitex_gen/like/likeservice"
	g "dousheng/pkg/global"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	tracing "github.com/kitex-contrib/tracer-opentracing"
)

var LikeClient likeservice.Client

func InitLikeRpc() {
	r, err := etcd.NewEtcdResolver([]string{g.EtcdAddress})
	if err != nil {
		panic(err)
	}
	cli, err := likeservice.NewClient(
		g.ServiceLikeName,
		client.WithSuite(tracing.NewDefaultClientSuite()), // tracer
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	LikeClient = cli
}
