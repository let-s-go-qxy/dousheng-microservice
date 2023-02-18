package api

import (
	"context"
	userService "dousheng/kitex_gen/user"
	"dousheng/pkg/etcd_discovery"
	g "dousheng/pkg/global"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	"strconv"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func UserInfo(c context.Context, ctx *app.RequestContext) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	myID, _ := strconv.Atoi(ctx.Query("user_id"))
	req := &userService.UserInfoRequest{
		UserId: int64(userId),
		MyId:   int64(myID),
	}
	resp, err := etcd_discovery.UserClient.UserInfo(c, req)
	if err != nil {
		ctx.JSON(consts.StatusOK, UserResponse{
			Response: Response{
				StatusCode: g.StatusCodeFail,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	userInfo := &User{}
	copier.Copy(userInfo, resp.User)
	ctx.JSON(consts.StatusOK, UserResponse{
		Response: Response{
			StatusCode: resp.GetStatusCode(),
			StatusMsg:  resp.GetStatusMsg(),
		}, User: *userInfo,
	})
}

func UserLogin(c context.Context, ctx *app.RequestContext) {
	req := &userService.UserLoginRequest{
		Username: ctx.Query("username"),
		Password: ctx.Query("password"),
	}
	resp, err := etcd_discovery.UserClient.UserLogin(c, req)
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
		UserId: resp.GetUserId(),
		Token:  resp.GetToken(),
	})
}

func UserRegister(c context.Context, ctx *app.RequestContext) {
	name := ctx.Query("username")
	pw := ctx.Query("password")
	req := userService.UserRegisterRequest{
		Username: name,
		Password: pw,
	}
	resp, err := etcd_discovery.UserClient.UserRegister(c, &req)
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
		UserId: resp.GetUserId(),
		Token:  resp.GetToken(),
	})
}
