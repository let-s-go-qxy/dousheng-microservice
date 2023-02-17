package main

import (
	feed "dousheng/kitex_gen/feed/feedservice"
	"dousheng/pkg/database"
	"dousheng/pkg/etcd_discovery"
	g "dousheng/pkg/global"
	"dousheng/pkg/tracer"
	"log"
)

func init() {
	tracer.InitJaeger(g.ServiceFeedName)
	etcd_discovery.InitUserRpc()
	database.InitDB()
}

func main() {
	svr := feed.NewServer(new(FeedServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
