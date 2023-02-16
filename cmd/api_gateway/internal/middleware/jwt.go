package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func Jwt() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		ctx.Next(c)
	}
}
