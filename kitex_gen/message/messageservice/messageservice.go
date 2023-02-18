// Code generated by Kitex v0.4.4. DO NOT EDIT.

package messageservice

import (
	"context"
	message "dousheng/kitex_gen/message"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return messageServiceServiceInfo
}

var messageServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MessageService"
	handlerType := (*message.MessageService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetMessageList":    kitex.NewMethodInfo(getMessageListHandler, newGetMessageListArgs, newGetMessageListResult, false),
		"PostMessageAction": kitex.NewMethodInfo(postMessageActionHandler, newPostMessageActionArgs, newPostMessageActionResult, false),
		"GetLatestMessage":  kitex.NewMethodInfo(getLatestMessageHandler, newGetLatestMessageArgs, newGetLatestMessageResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "message",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func getMessageListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(message.MessageChatRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(message.MessageService).GetMessageList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetMessageListArgs:
		success, err := handler.(message.MessageService).GetMessageList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetMessageListResult)
		realResult.Success = success
	}
	return nil
}
func newGetMessageListArgs() interface{} {
	return &GetMessageListArgs{}
}

func newGetMessageListResult() interface{} {
	return &GetMessageListResult{}
}

type GetMessageListArgs struct {
	Req *message.MessageChatRequest
}

func (p *GetMessageListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(message.MessageChatRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetMessageListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetMessageListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetMessageListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetMessageListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetMessageListArgs) Unmarshal(in []byte) error {
	msg := new(message.MessageChatRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetMessageListArgs_Req_DEFAULT *message.MessageChatRequest

func (p *GetMessageListArgs) GetReq() *message.MessageChatRequest {
	if !p.IsSetReq() {
		return GetMessageListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetMessageListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetMessageListResult struct {
	Success *message.MessageChatResponse
}

var GetMessageListResult_Success_DEFAULT *message.MessageChatResponse

func (p *GetMessageListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(message.MessageChatResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetMessageListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetMessageListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetMessageListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetMessageListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetMessageListResult) Unmarshal(in []byte) error {
	msg := new(message.MessageChatResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMessageListResult) GetSuccess() *message.MessageChatResponse {
	if !p.IsSetSuccess() {
		return GetMessageListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetMessageListResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.MessageChatResponse)
}

func (p *GetMessageListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func postMessageActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(message.MessageActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(message.MessageService).PostMessageAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PostMessageActionArgs:
		success, err := handler.(message.MessageService).PostMessageAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PostMessageActionResult)
		realResult.Success = success
	}
	return nil
}
func newPostMessageActionArgs() interface{} {
	return &PostMessageActionArgs{}
}

func newPostMessageActionResult() interface{} {
	return &PostMessageActionResult{}
}

type PostMessageActionArgs struct {
	Req *message.MessageActionRequest
}

func (p *PostMessageActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(message.MessageActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PostMessageActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PostMessageActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PostMessageActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PostMessageActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PostMessageActionArgs) Unmarshal(in []byte) error {
	msg := new(message.MessageActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PostMessageActionArgs_Req_DEFAULT *message.MessageActionRequest

func (p *PostMessageActionArgs) GetReq() *message.MessageActionRequest {
	if !p.IsSetReq() {
		return PostMessageActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PostMessageActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type PostMessageActionResult struct {
	Success *message.MessageActionResponse
}

var PostMessageActionResult_Success_DEFAULT *message.MessageActionResponse

func (p *PostMessageActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(message.MessageActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PostMessageActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PostMessageActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PostMessageActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PostMessageActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PostMessageActionResult) Unmarshal(in []byte) error {
	msg := new(message.MessageActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PostMessageActionResult) GetSuccess() *message.MessageActionResponse {
	if !p.IsSetSuccess() {
		return PostMessageActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PostMessageActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.MessageActionResponse)
}

func (p *PostMessageActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getLatestMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(message.MessageLastRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(message.MessageService).GetLatestMessage(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetLatestMessageArgs:
		success, err := handler.(message.MessageService).GetLatestMessage(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetLatestMessageResult)
		realResult.Success = success
	}
	return nil
}
func newGetLatestMessageArgs() interface{} {
	return &GetLatestMessageArgs{}
}

func newGetLatestMessageResult() interface{} {
	return &GetLatestMessageResult{}
}

type GetLatestMessageArgs struct {
	Req *message.MessageLastRequest
}

func (p *GetLatestMessageArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(message.MessageLastRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetLatestMessageArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetLatestMessageArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetLatestMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetLatestMessageArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetLatestMessageArgs) Unmarshal(in []byte) error {
	msg := new(message.MessageLastRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetLatestMessageArgs_Req_DEFAULT *message.MessageLastRequest

func (p *GetLatestMessageArgs) GetReq() *message.MessageLastRequest {
	if !p.IsSetReq() {
		return GetLatestMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetLatestMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetLatestMessageResult struct {
	Success *message.MessageLastResponse
}

var GetLatestMessageResult_Success_DEFAULT *message.MessageLastResponse

func (p *GetLatestMessageResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(message.MessageLastResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetLatestMessageResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetLatestMessageResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetLatestMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetLatestMessageResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetLatestMessageResult) Unmarshal(in []byte) error {
	msg := new(message.MessageLastResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLatestMessageResult) GetSuccess() *message.MessageLastResponse {
	if !p.IsSetSuccess() {
		return GetLatestMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetLatestMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.MessageLastResponse)
}

func (p *GetLatestMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetMessageList(ctx context.Context, Req *message.MessageChatRequest) (r *message.MessageChatResponse, err error) {
	var _args GetMessageListArgs
	_args.Req = Req
	var _result GetMessageListResult
	if err = p.c.Call(ctx, "GetMessageList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PostMessageAction(ctx context.Context, Req *message.MessageActionRequest) (r *message.MessageActionResponse, err error) {
	var _args PostMessageActionArgs
	_args.Req = Req
	var _result PostMessageActionResult
	if err = p.c.Call(ctx, "PostMessageAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetLatestMessage(ctx context.Context, Req *message.MessageLastRequest) (r *message.MessageLastResponse, err error) {
	var _args GetLatestMessageArgs
	_args.Req = Req
	var _result GetLatestMessageResult
	if err = p.c.Call(ctx, "GetLatestMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
