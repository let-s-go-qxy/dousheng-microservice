package main

import (
	video "dousheng/kitex_gen/video/videoservice"
	"dousheng/pkg/database"
	"dousheng/pkg/etcd_discovery"
	g "dousheng/pkg/global"
	"dousheng/pkg/tracer"
	"log"
)

func init() {
	tracer.InitJaeger(g.ServiceVideoName)
	etcd_discovery.InitUserRpc()
	database.InitDB()
}

func main() {
	svr := video.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
