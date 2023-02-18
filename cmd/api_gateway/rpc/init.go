package rpc

import (
	"dousheng/pkg/etcd_discovery"
)

// InitRPC init rpc client
func InitRPC() {
	etcd_discovery.InitUserRpc()
	etcd_discovery.InitVideoRpc()
	etcd_discovery.InitRelationRpc()
	etcd_discovery.InitMessageRpc()
	etcd_discovery.InitLikeRpc()
	etcd_discovery.InitVideoRpc()
}
