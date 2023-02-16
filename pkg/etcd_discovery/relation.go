package etcd_discovery

import (
	"dousheng/kitex_gen/relation/relationservice"
	g "dousheng/pkg/global"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	tracing "github.com/kitex-contrib/tracer-opentracing"
)

var RelationClient relationservice.Client

func InitRelationRpc() {
	r, err := etcd.NewEtcdResolver([]string{g.EtcdAddress})
	if err != nil {
		panic(err)
	}
	cli, err := relationservice.NewClient(
		g.ServiceRelationName,
		client.WithSuite(tracing.NewDefaultClientSuite()), // tracer
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	RelationClient = cli
}
