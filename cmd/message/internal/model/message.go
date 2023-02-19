package model

import (
	g "dousheng/pkg/global"
	"dousheng/pkg/mq"
	m "dousheng/pkg/mq"
	"github.com/cloudwego/kitex/pkg/klog"
)

type Message struct {
	Id         int    `json:"id"`
	FromId     int64  `json:"from_id"`
	ToId       int64  `json:"to_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

type MessageSendEvent struct {
	UserId     int    `json:"user_id,"`
	ToUserId   int    `json:"to_user_id"`
	MsgContent string `json:"msg_content"`
}

type MessagePushEvent struct {
	FromUserId int    `json:"user_id"`
	MsgContent string `json:"msg_content"`
}

type RespMessage struct {
	Id         int    `json:"id"`
	ToId       int    `json:"to_id"`
	FromId     int    `json:"from_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

func GetMessageListMQ(fromID int) ([]mq.RespMessage, error) {
	recordListMQ, err := m.GetRabbitMQMessageList(fromID)
	if err != nil {
		klog.Error("GetRabbitMQMessageList时出错了:" + err.Error())
	}
	currentListMQ, err := m.GetRabbitMQMessageCurrent(fromID)
	if err != nil {
		klog.Error("GetRabbitMQMessageCurrent时出错了:" + err.Error())
	}
	allMessageListMQ := append(recordListMQ, currentListMQ...)
	return allMessageListMQ, err
}

// CreateMessage 创建消息
func CreateMessage(message *RespMessage) (err error) {
	err = g.MysqlDB.Table("messages").Create(message).Error
	return
}

// GetMsgLatest 获取最新的聊天记录  msgType 0为接收的信息，1为发送的信息
func GetMsgLatest(userId, myId int) (msg string, msgType int) {
	msgDao := new(RespMessage)
	g.MysqlDB.Table("messages").Where("to_id = ? AND from_id = ?", userId, myId).
		Or("to_id = ? AND from_id = ?", myId, userId).Order("create_time desc").First(msgDao)
	if msgDao.ToId == userId {
		msgType = 1
	}
	msg = msgDao.Content
	return
}

func GetMsgListByDB(fromID, toID int64) ([]Message, error) {
	messages := make([]Message, 0)
	err := g.MysqlDB.Order("create_time").Find(&messages, "from_id = ? AND to_id = ?", fromID, toID).Error
	return messages, err
}
