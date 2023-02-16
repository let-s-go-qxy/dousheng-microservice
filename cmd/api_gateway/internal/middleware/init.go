package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
)

// InitMiddleWareForDefault 此处配置全局的middleware
func InitMiddleWareForDefault(h *server.Hertz) *server.Hertz {
	h.Use(Cors(), AccessLog())
	return h
}

// AccessLog hertz log
func AccessLog() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()
		ctx.Next(c)
		end := time.Now()
		latency := end.Sub(start).Milliseconds()
		hlog.CtxTracef(c, " | %s | %s    ==>  [%d] --- %d ms",
			ctx.Request.Header.Method(), ctx.Request.URI().PathOriginal(), ctx.Response.StatusCode(), latency)
	}
}
