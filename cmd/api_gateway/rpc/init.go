package rpc

import (
	"dousheng/pkg/etcd_discovery"
)

// InitRPC init rpc client
func InitRPC() {
	etcd_discovery.InitUserRpc()
	etcd_discovery.InitVideoRpc()
	etcd_discovery.InitRelationRpc()
}
