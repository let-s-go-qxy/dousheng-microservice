package middleware

import (
	"context"
	"dousheng/cmd/api_gateway/api"
	"dousheng/cmd/api_gateway/internal/service/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"time"
)

func Jwt() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		token := ctx.Query("token")
		if token2 := ctx.PostForm("token"); token2 != "" {
			token = token2
		}
		if token == "" {
			ctx.AbortWithStatusJSON(consts.StatusOK, api.Response{
				StatusCode: 1, StatusMsg: "User doesn't exist",
			})
			return
		}
		claims, err := auth.ParseToken(token)
		if err != nil || claims.Expire < int(time.Now().Unix()) {
			// TODO 这里应该是返回401，但是Demo中是这么写的，所以为了适配app故此
			ctx.AbortWithStatusJSON(consts.StatusOK, api.Response{
				StatusCode: 1, StatusMsg: "User doesn't exist",
			})
			return
		}
		ctx.Set("user_id", claims.Id)
		ctx.Next(c)
	}
}
