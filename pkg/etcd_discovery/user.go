package etcd_discovery

import (
	"context"
	userService "dousheng/kitex_gen/user"
	"dousheng/kitex_gen/user/userservice"
	g "dousheng/pkg/global"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	tracing "github.com/kitex-contrib/tracer-opentracing"
	"github.com/pkg/errors"
)

var userClient userservice.Client

// InitUserRpc 服务发现
func InitUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{g.EtcdAddress})
	if err != nil {
		panic(err)
	}
	cli, err := userservice.NewClient(
		g.ServiceUserName,
		client.WithSuite(tracing.NewDefaultClientSuite()), // tracer
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	userClient = cli
}

// UserInfo 定义响应的微服务业务方法
func UserInfo(ctx context.Context, req *userService.UserInfoRequest) (*userService.UserInfoResponse, error) {
	resp, err := userClient.UserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New("resp error")
	}
	return resp, nil
}
