package etcd_discovery

import (
	"context"
	relationService "dousheng/kitex_gen/relation"
	"dousheng/kitex_gen/relation/relationservice"
	g "dousheng/pkg/global"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	tracing "github.com/kitex-contrib/tracer-opentracing"
	"github.com/pkg/errors"
)

var relationClient relationservice.Client

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
	relationClient = cli
}

func FollowerList(ctx context.Context, req *relationService.RelationFollowListRequest) (*relationService.RelationFollowListResponse, error) {
	resp, err := relationClient.RelationFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New("resp error")
	}
	return resp, nil
}
