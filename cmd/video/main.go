package main

import (
	video "dousheng/kitex_gen/video/videoservice"
	"log"
)

func main() {
	svr := video.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
