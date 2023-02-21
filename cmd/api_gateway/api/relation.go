package api

import (
	"context"
	"dousheng/kitex_gen/relation"
	"dousheng/pkg/etcd_discovery"
	"dousheng/pkg/mq"
	utils2 "dousheng/pkg/utils"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
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

// RelationAction 关注
func RelationAction(c context.Context, ctx *app.RequestContext) {
	toUserId, _ := strconv.Atoi(ctx.Query("to_user_id"))
	actionType, _ := strconv.Atoi(ctx.Query("action_type"))
	myId, _ := ctx.Get("user_id")
	resp, err := etcd_discovery.RelationClient.RelationAction(c, &relation.RelationActionRequest{
		FromUserId: myId.(int64),
		ToUserId:   int64(toUserId),
		ActionType: int32(actionType),
	})
	if err != nil {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	ctx.JSON(consts.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  resp.StatusMsg,
	})
}

// GetFollowList 获取关注列表
func GetFollowList(c context.Context, ctx *app.RequestContext) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	myId, _ := ctx.Get("user_id")
	resp, err := etcd_discovery.RelationClient.GetFollowList(c, &relation.RelationFollowListRequest{
		UserId: int64(userId),
		MyId:   myId.(int64),
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
	// 防止传null
	if len(resp.GetUserList()) == 0 {
		ctx.JSON(consts.StatusOK, utils.H{
			"status_code": 0,
			"status_msg":  "ok",
			"user_list":   []User{},
		})
		return
	}
	ctx.JSON(consts.StatusOK, utils2.ConvertStruct(resp, nil))
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(c context.Context, ctx *app.RequestContext) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	myId, _ := ctx.Get("user_id")
	resp, err := etcd_discovery.RelationClient.GetFollowerList(c, &relation.RelationFollowerListRequest{
		UserId: int64(userId),
		MyId:   myId.(int64),
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
	// 防止传null
	if len(resp.GetUserList()) == 0 {
		ctx.JSON(consts.StatusOK, utils.H{
			"status_code": 0,
			"status_msg":  "ok",
			"user_list":   []User{},
		})
		return
	}
	ctx.JSON(consts.StatusOK, utils2.ConvertStruct(resp, nil))
}

// GetFriendList 获取好友列表 同时获取最新的聊天记录
func GetFriendList(c context.Context, ctx *app.RequestContext) {
	myId, _ := ctx.Get("user_id")
	resp, err := etcd_discovery.RelationClient.GetFriendList(c, &relation.RelationFriendListRequest{
		UserId: myId.(int64),
		MyId:   myId.(int64),
	})

	allFriendsMessageListRsp, _ := etcd_discovery.RelationClient.GetFriendsMessageList(c, &relation.RelationFriendsMessageListRequest{
		UserId: myId.(int64),
	})
	var list []RespMessage
	for _, m := range allFriendsMessageListRsp.GetMessageList() {

		t := time.Time{}
		t, _ = time.ParseInLocation("2006-01-02T15:04:05Z07:00", m.GetCreateTime(), time.Local)

		list = append(list, RespMessage{
			Id:         int(m.GetId()),
			ToId:       int(m.GetToId()),
			FromId:     int(m.GetFromId()),
			Content:    m.GetContent(),
			CreateTime: strconv.Itoa(int(t.Unix())),
		})
	}
	marshal, _ := json.Marshal(list)
	//将allFriendsMessageList所有的朋友的聊天记录放到消息队列中去
	strJson := string(marshal)
	err = mq.PublishMessageListToMQ(strJson, int(myId.(int64)))
	if err != nil {
		ctx.JSON(consts.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	// 防止传null
	if len(resp.GetUserList()) == 0 {
		ctx.JSON(consts.StatusOK, utils.H{
			"status_code": 0,
			"status_msg":  "ok",
			"user_list":   []User{},
		})
		return
	}
	ctx.JSON(consts.StatusOK, utils2.ConvertStruct(resp, nil))
}
