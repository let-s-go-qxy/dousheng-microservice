package api

import (
	"context"
	userService "dousheng/kitex_gen/user"
	"dousheng/pkg/etcd_discovery"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type UserLoginResponse struct {
	Response
	UserId int    `json:"user_id,omitempty"`
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
	resp, err := etcd_discovery.UserInfo(context.Background(), req)
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
