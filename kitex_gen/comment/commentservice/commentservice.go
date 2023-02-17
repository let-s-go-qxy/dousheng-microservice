// Code generated by Kitex v0.4.4. DO NOT EDIT.

package commentservice

import (
	"context"
	comment "dousheng/kitex_gen/comment"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

var commentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CommentService"
	handlerType := (*comment.CommentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"PostCommentAction": kitex.NewMethodInfo(postCommentActionHandler, newPostCommentActionArgs, newPostCommentActionResult, false),
		"GetCommentList":    kitex.NewMethodInfo(getCommentListHandler, newGetCommentListArgs, newGetCommentListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "comment",
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

func postCommentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.CommentActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).PostCommentAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PostCommentActionArgs:
		success, err := handler.(comment.CommentService).PostCommentAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PostCommentActionResult)
		realResult.Success = success
	}
	return nil
}
func newPostCommentActionArgs() interface{} {
	return &PostCommentActionArgs{}
}

func newPostCommentActionResult() interface{} {
	return &PostCommentActionResult{}
}

type PostCommentActionArgs struct {
	Req *comment.CommentActionRequest
}

func (p *PostCommentActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.CommentActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PostCommentActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PostCommentActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PostCommentActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PostCommentActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PostCommentActionArgs) Unmarshal(in []byte) error {
	msg := new(comment.CommentActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PostCommentActionArgs_Req_DEFAULT *comment.CommentActionRequest

func (p *PostCommentActionArgs) GetReq() *comment.CommentActionRequest {
	if !p.IsSetReq() {
		return PostCommentActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PostCommentActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type PostCommentActionResult struct {
	Success *comment.CommentActionResponse
}

var PostCommentActionResult_Success_DEFAULT *comment.CommentActionResponse

func (p *PostCommentActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.CommentActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PostCommentActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PostCommentActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PostCommentActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PostCommentActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PostCommentActionResult) Unmarshal(in []byte) error {
	msg := new(comment.CommentActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PostCommentActionResult) GetSuccess() *comment.CommentActionResponse {
	if !p.IsSetSuccess() {
		return PostCommentActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PostCommentActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.CommentActionResponse)
}

func (p *PostCommentActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getCommentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.CommentListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).GetCommentList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetCommentListArgs:
		success, err := handler.(comment.CommentService).GetCommentList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetCommentListResult)
		realResult.Success = success
	}
	return nil
}
func newGetCommentListArgs() interface{} {
	return &GetCommentListArgs{}
}

func newGetCommentListResult() interface{} {
	return &GetCommentListResult{}
}

type GetCommentListArgs struct {
	Req *comment.CommentListRequest
}

func (p *GetCommentListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.CommentListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetCommentListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetCommentListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetCommentListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetCommentListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetCommentListArgs) Unmarshal(in []byte) error {
	msg := new(comment.CommentListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetCommentListArgs_Req_DEFAULT *comment.CommentListRequest

func (p *GetCommentListArgs) GetReq() *comment.CommentListRequest {
	if !p.IsSetReq() {
		return GetCommentListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCommentListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetCommentListResult struct {
	Success *comment.CommentListResponse
}

var GetCommentListResult_Success_DEFAULT *comment.CommentListResponse

func (p *GetCommentListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.CommentListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetCommentListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetCommentListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetCommentListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetCommentListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetCommentListResult) Unmarshal(in []byte) error {
	msg := new(comment.CommentListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCommentListResult) GetSuccess() *comment.CommentListResponse {
	if !p.IsSetSuccess() {
		return GetCommentListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCommentListResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.CommentListResponse)
}

func (p *GetCommentListResult) IsSetSuccess() bool {
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

func (p *kClient) PostCommentAction(ctx context.Context, Req *comment.CommentActionRequest) (r *comment.CommentActionResponse, err error) {
	var _args PostCommentActionArgs
	_args.Req = Req
	var _result PostCommentActionResult
	if err = p.c.Call(ctx, "PostCommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCommentList(ctx context.Context, Req *comment.CommentListRequest) (r *comment.CommentListResponse, err error) {
	var _args GetCommentListArgs
	_args.Req = Req
	var _result GetCommentListResult
	if err = p.c.Call(ctx, "GetCommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}