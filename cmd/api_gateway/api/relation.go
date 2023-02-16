package api

import (
	"context"
	"dousheng/kitex_gen/relation"
	"dousheng/pkg/etcd_discovery"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	"time"
)

type UserListResponse struct {
	Response
	UserList []interface{} `json:"user_list"`
}

type UserAndMsg struct {
	User
	Message string `json:"message"`
	MsgType int    `json:"msg_type"`
}

type UserAndMsgListResponse struct {
	Response
	UserList []UserAndMsg `json:"user_list"`
}

func GetFollowList(c context.Context, ctx *app.RequestContext) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	resp, err := etcd_discovery.RelationClient.RelationFollowList(c, &relation.RelationFollowListRequest{
		UserId: int64(userId),
		Token:  "",
	})
	if err != nil {
		ctx.JSON(consts.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	time.Sleep(2 * time.Second)
	ctx.JSON(consts.StatusOK, resp)
}
