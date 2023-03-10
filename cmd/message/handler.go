package main

import (
	"context"
	messageService "dousheng/cmd/message/internal/service"
	message "dousheng/kitex_gen/message"
	g "dousheng/pkg/global"
	"github.com/jinzhu/copier"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// GetMessageList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessageList(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	allMessageListMQ, err := messageService.GetMessageList(int(req.UserId))
	allMessageList := []*message.Message{}
	//copier.Copy(&allMessageList, &allMessageListMQ)

	for _, respMessage := range allMessageListMQ {
		item := message.Message{}
		copier.Copy(&item, &respMessage)
		item.Id = int64(respMessage.Id)
		item.ToId = int64(respMessage.ToId)
		item.FromId = int64(respMessage.FromId)
		allMessageList = append(allMessageList, &item)
	}
	response := &message.MessageChatResponse{}
	response.StatusCode = g.StatusCodeOk
	response.StatusMsg = "获取聊天记录成功"
	response.MessageList = allMessageList
	return response, err
}

// PostMessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) PostMessageAction(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {

	err = messageService.PostMessageAction(int(req.UserId), int(req.ToUserId), req.Content, int(req.ActionType))
	if err != nil {
		messageResponse := &message.MessageActionResponse{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  "发送消息失败",
		}
		return messageResponse, err
	}
	messageResponse := &message.MessageActionResponse{
		StatusCode: g.StatusCodeOk,
		StatusMsg:  "发送消息成功",
	}

	return messageResponse, err
}

// GetMessageListByDB implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessageListByDB(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	return messageService.GetMessageListByDB(ctx, req.GetUserId(), req.GetToUserId())
}

// GetLatestMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetLatestMessage(ctx context.Context, req *message.MessageLatestRequest) (resp *message.MessageLatestResponse, err error) {
	msg, msgType := messageService.GetMsgLatest(int(req.GetUserId()), int(req.GetMyId()))
	resp = &message.MessageLatestResponse{
		Content: msg,
		MsgType: int32(msgType),
	}
	return
}
