package main

import (
	"context"
	messageService "dousheng/cmd/message/internal/service"
	message "dousheng/kitex_gen/message"
	g "dousheng/pkg/global"
	m "dousheng/pkg/mq"
	"github.com/jinzhu/copier"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// GetMessageList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessageList(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	current, err := m.GetRabbitMQMessageCurrent(int(req.UserId))
	list, err := m.GetRabbitMQMessageList(int(req.UserId))
	allMessageListMQ := append(current, list...)
	allMessageList := []*message.Message{}
	copier.Copy(&allMessageList, &allMessageListMQ)

	for _, respMessage := range allMessageListMQ {
		item := message.Message{}
		copier.Copy(&item, &respMessage)
		allMessageList = append(allMessageList, &item)
	}
	response := &message.MessageChatResponse{
		StatusCode:  g.StatusCodeOk,
		StatusMsg:   "获取聊天记录成功",
		MessageList: allMessageList,
	}
	return response, err
}

// PostMessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) PostMessageAction(ctx context.Context, req *message.RelationActionRequest) (resp *message.RelationActionResponse, err error) {

	err = messageService.MessageAction(int(req.UserId), int(req.ToUserId), req.Content, int(req.ActionType))
	if err != nil {
		messageResponse := &message.RelationActionResponse{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  "发送消息失败",
		}
		return messageResponse, err
	}
	messageResponse := &message.RelationActionResponse{
		StatusCode: g.StatusCodeOk,
		StatusMsg:  "发送消息成功",
	}

	return messageResponse, err
}
