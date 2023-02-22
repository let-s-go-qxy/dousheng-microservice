package main

import (
	"context"
	"dousheng/cmd/api_gateway/router"
	"dousheng/cmd/api_gateway/rpc"
	"dousheng/pkg/database"
	g "dousheng/pkg/global"
	"dousheng/pkg/tracer"
	"flag"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/pprof"
	"log"
	"time"
)

func init() {
	tracer.InitJaeger(g.ApiGatewayName)
	rpc.InitRPC()
	database.InitDB()
}

// pprof
func counter() {
	slice := make([]int, 0)
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
	}
}

// pprof
func workForever() {
	for {
		go counter()
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go workForever() // pprof
	flag.Parse()
	readTimeout, err := time.ParseDuration("1m")
	if err != nil {
		fmt.Println("parse duration err")
	}
	writeTimout, err := time.ParseDuration("1m")
	if err != nil {
		log.Println("parse duration err")
	}
	h := server.New(
		server.WithHostPorts(g.ApiGatewayAddress),
		server.WithReadTimeout(readTimeout),
		server.WithWriteTimeout(writeTimout),
		server.WithHandleMethodNotAllowed(true), // 不加会出现 jaeger 失效
	)
	// 熔断
	h.Use(recovery.Recovery(recovery.WithRecoveryHandler(
		func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
			hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
			c.JSON(consts.StatusInternalServerError, map[string]interface{}{
				"code":    1,
				"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
			})
		})))
	router.InitRouter(h)
	h.NoRoute(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no route")
	})
	h.NoMethod(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no method")
	})
	pprof.Register(h)
	h.Spin()
}
