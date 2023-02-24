package service

import (
	"context"
	"dousheng/cmd/message/internal/model"
	"dousheng/cmd/message/num"
	"dousheng/kitex_gen/message"
	"dousheng/pkg/mq"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"sync/atomic"
	"time"
)

// GetMsgLatest 获取最新的聊天记录  msgType 0为接收的信息，1为发送的信息
func GetMsgLatest(userId, myId int) (msg string, msgType int) {
	msg, msgType = model.GetMsgLatest(userId, myId)
	return
}

// GetMessageList 从MQ中获取队列中的消息进行消费
func GetMessageList(fromID int) ([]mq.RespMessage, error) {
	listMQ, err := model.GetMessageListMQ(fromID)
	if err != nil {
		klog.Error(err.Error())
	}
	return listMQ, err
}

func PostMessageActionWithMQ(fromId int, toId int, content string, actionType int) (err error) {
	atomic.AddInt32(&num.ReqNUM, 1)
	marshal, _ := json.Marshal(model.RespMessage{
		ToId:       toId,
		FromId:     fromId,
		Content:    content,
		CreateTime: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	})
	if err != nil {
		return err
	}
	err = mq.SendSyncMessage(string(marshal))
	if err != nil {
		return err
	}
	return nil
}

func SubscribeMessageAction(c context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	data := model.RespMessage{}
	for i := range msgs {
		atomic.AddInt32(&num.MessageNum, 1)
		json.Unmarshal(msgs[i].Message.Body, &data)
		err := PostMessageAction(data.FromId, data.ToId, data.Content, 1)
		if err != nil {
			return consumer.ConsumeRetryLater, err
		}
	}
	return consumer.ConsumeSuccess, nil
}

func PostMessageAction(fromId int, toId int, content string, actionType int) (err error) {

	msg := model.Message{
		ToId:       int64(toId),
		FromId:     int64(fromId),
		Content:    content,
		CreateTime: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	return model.PostMessageOne(msg)
	//JsonMsg, err := json.Marshal(msg)
	//strJsonMsg := string(JsonMsg)

	//将消息写到userId对应的消息队列中去
	//mq.PublishMessageCurrentToMQ(strJsonMsg, fromId)
	//将消息写到ToId对应的消息队列中去
	//mq.PublishMessageCurrentToMQ(strJsonMsg, toId)
}

func GetMessageListByDB(c context.Context, fromId, toId int64) (resp *message.MessageChatResponse, err error) {
	resp = &message.MessageChatResponse{}
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
