package api

import (
	"context"
	userService "dousheng/kitex_gen/user"
	"dousheng/pkg/etcd_discovery"
	g "dousheng/pkg/global"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func UserInfo(c context.Context, ctx *app.RequestContext) {
	req := &userService.UserInfoRequest{
		UserId: 1,
		Token:  "123",
	}
	resp, err := etcd_discovery.UserClient.UserInfo(c, req)
	if err != nil {
		ctx.JSON(consts.StatusOK, UserResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	ctx.JSON(consts.StatusOK, resp)
}

func UserLogin(c context.Context, ctx *app.RequestContext) {
	req := &userService.UserLoginRequest{
		Username: ctx.Query("username"),
		Password: ctx.Query("password"),
	}
	userInfo, err := etcd_discovery.UserClient.UserLogin(c, req)
	if err != nil {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  err.Error(),
		})
		return
	}
	ctx.JSON(consts.StatusOK, UserLoginResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserId: userInfo.UserId,
		Token:  userInfo.Token,
	})
}
