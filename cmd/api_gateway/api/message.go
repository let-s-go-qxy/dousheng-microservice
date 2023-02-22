package api

import (
	"context"
	"dousheng/kitex_gen/message"
	"dousheng/pkg/etcd_discovery"
	g "dousheng/pkg/global"
	utils2 "dousheng/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	"strconv"
)

type RespMessage struct {
	Id         int    `json:"id"`
	ToId       int    `json:"to_id"`
	FromId     int    `json:"from_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

type MergeMessage struct {
	Id         int    `json:"id"`
	ToId       int    `json:"to_user_id"`
	FromId     int    `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}
type MessageListResponse struct {
	Response
	MessageList []MergeMessage `json:"message_list"`
}

func GetMessageList(c context.Context, ctx *app.RequestContext) {

	//获取to_user_id和user_id

	toId, err := strconv.Atoi(ctx.Query("to_user_id"))
	if err != nil {

		//g.Logger.Error("获取对方ID错误")
	}

	userIdInterface, success := ctx.Get("user_id")
	var fromId int
	if success {
		fromId = int(userIdInterface.(int64))
	} // 若不存在，userID默认为0
	messageChatRequest := &message.MessageChatRequest{}
	messageChatRequest.UserId = int64(fromId)
	messageChatRequest.ToUserId = int64(toId)

	getMessageListResponse, err := etcd_discovery.MessageClient.GetMessageList(c, messageChatRequest)
	messages := getMessageListResponse.GetMessageList()
	var msg RespMessage
	allMessageList := []RespMessage{}
	for _, msgPointer := range messages {
		m := *msgPointer
		copier.Copy(&msg, &m)
		msg.ToId = int(m.ToId)
		msg.FromId = int(m.FromId)
		allMessageList = append(allMessageList, msg)
	}
	messageList := []RespMessage{}
	for _, message := range allMessageList {
		if (message.ToId == toId && message.FromId == fromId) || (message.ToId == fromId && message.FromId == toId) {
			messageList = append(messageList, message)
		}
	}
	if err != nil {
		hlog.Error("GetRabbitMQMessageList时发生了错误！")
	}
	//messageList, _ := m.GetMessageList(toId, fromId)
	respMessageList := make([]MergeMessage, 0)
	copier.Copy(&respMessageList, &messageList)
	resp := MessageListResponse{Response: Response{
		StatusCode: g.StatusCodeOk,
		StatusMsg:  "获取消息列表成功!!"},
		MessageList: respMessageList}

	//marshal, _ := json.Marshal(respMessageList)
	//fmt.Println(string(marshal))
	ctx.JSON(consts.StatusOK, utils2.ConvertStruct(resp, nil))

}

func PostMessageAction(c context.Context, ctx *app.RequestContext) {

	userIDInterface, success := ctx.Get("user_id")
	var fromId int
	if success {
		fromId = int(userIDInterface.(int64))
	} // 若不存在，userID默认为0

	toId, _ := strconv.Atoi(ctx.Query("to_user_id"))
	content := ctx.Query("content")
	actionType, _ := strconv.Atoi(ctx.Query("action_type"))
	chatActionRequest := &message.MessageActionRequest{
		UserId:     int64(fromId),
		ToUserId:   int64(toId),
		ActionType: int32(actionType),
		Content:    content,
	}

	_, err := etcd_discovery.MessageClient.PostMessageAction(c, chatActionRequest)

	if err != nil {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  err.Error(),
		})
	}

	resp := Response{StatusCode: g.StatusCodeOk, StatusMsg: "发送消息成功"}
	ctx.JSON(consts.StatusOK, resp)

}
