// Code generated by Kitex v0.4.4. DO NOT EDIT.

package messageservice

import (
	"context"
	message "dousheng/kitex_gen/message"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetMessageList(ctx context.Context, Req *message.MessageChatRequest, callOptions ...callopt.Option) (r *message.MessageChatResponse, err error)
	GetMessageListByDB(ctx context.Context, Req *message.MessageChatRequest, callOptions ...callopt.Option) (r *message.MessageChatResponse, err error)
	PostMessageAction(ctx context.Context, Req *message.MessageActionRequest, callOptions ...callopt.Option) (r *message.MessageActionResponse, err error)
	GetLatestMessage(ctx context.Context, Req *message.MessageLastRequest, callOptions ...callopt.Option) (r *message.MessageLastResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kMessageServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kMessageServiceClient struct {
	*kClient
}

func (p *kMessageServiceClient) GetMessageList(ctx context.Context, Req *message.MessageChatRequest, callOptions ...callopt.Option) (r *message.MessageChatResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMessageList(ctx, Req)
}

func (p *kMessageServiceClient) GetMessageListByDB(ctx context.Context, Req *message.MessageChatRequest, callOptions ...callopt.Option) (r *message.MessageChatResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMessageListByDB(ctx, Req)
}

func (p *kMessageServiceClient) PostMessageAction(ctx context.Context, Req *message.MessageActionRequest, callOptions ...callopt.Option) (r *message.MessageActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PostMessageAction(ctx, Req)
}

func (p *kMessageServiceClient) GetLatestMessage(ctx context.Context, Req *message.MessageLastRequest, callOptions ...callopt.Option) (r *message.MessageLastResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLatestMessage(ctx, Req)
}
