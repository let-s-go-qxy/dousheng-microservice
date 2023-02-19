package service

import (
	"context"
	"dousheng/cmd/message/internal/model"
	"dousheng/kitex_gen/message"
	g "dousheng/pkg/global"
	"dousheng/pkg/mq"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"time"
)

func GetMessageList(fromID int) ([]mq.RespMessage, error) {
	listMQ, err := model.GetMessageListMQ(fromID)
	if err != nil {
		klog.Error(err.Error())
	}
	return listMQ, err
}
func PostMessageAction(fromId int, toId int, content string, actionType int) (err error) {

	msg := model.RespMessage{
		ToId:       toId,
		FromId:     fromId,
		Content:    content,
		CreateTime: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	if actionType == g.MessageSendEvent {
		err = model.CreateMessage(&msg)
		if err != nil {
			//err = errors.New("发送消息失败: " + err.Error())
		}
	}
	fmt.Println(msg.CreateTime)

	t := time.Time{}
	//fmt.Println(message.CreateTime)
	t, _ = time.ParseInLocation("2006-01-02T15:04:05Z07:00", msg.CreateTime, time.Local)
	msg.CreateTime = strconv.Itoa(int(t.Unix()))

	JsonMsg, err := json.Marshal(msg)
	strJsonMsg := string(JsonMsg)

	//将消息写到userId对应的消息队列中去
	//mq.PublishMessageCurrentToMQ(strJsonMsg, fromId)
	//将消息写到ToId对应的消息队列中去
	mq.PublishMessageCurrentToMQ(strJsonMsg, toId)

	return
}

func GetMessageListByDB(c context.Context, fromId, toId int64) (resp *message.MessageChatResponse, err error) {
	list, err := model.GetMsgListByDB(fromId, toId)
	if err != nil {
		return nil, err
	}
	resp.MessageList = make([]*message.Message, 0)
	for _, m := range list {
		resp.MessageList = append(resp.MessageList, &message.Message{
			Id:         int64(m.Id),
			ToId:       m.ToId,
			FromId:     m.FromId,
			Content:    m.Content,
			CreateTime: m.CreateTime,
		})
	}
	return
}
