package api

import (
	"context"
	"dousheng/kitex_gen/like"
	"dousheng/pkg/etcd_discovery"
)

func RefreshLikeCache() {
	req := &like.RefreshLikeCacheRequest{}
	etcd_discovery.LikeClient.RefreshLikeCache(context.Background(), req)

}
