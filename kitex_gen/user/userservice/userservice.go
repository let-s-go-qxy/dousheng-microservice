// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	user "dousheng/kitex_gen/user"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"UserInfo":           kitex.NewMethodInfo(userInfoHandler, newUserInfoArgs, newUserInfoResult, false),
		"UserLogin":          kitex.NewMethodInfo(userLoginHandler, newUserLoginArgs, newUserLoginResult, false),
		"UserRegister":       kitex.NewMethodInfo(userRegisterHandler, newUserRegisterArgs, newUserRegisterResult, false),
		"GetAvatar":          kitex.NewMethodInfo(getAvatarHandler, newGetAvatarArgs, newGetAvatarResult, false),
		"GetBackgroundImage": kitex.NewMethodInfo(getBackgroundImageHandler, newGetBackgroundImageArgs, newGetBackgroundImageResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
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

func userInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UserInfoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UserInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserInfoArgs:
		success, err := handler.(user.UserService).UserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserInfoResult)
		realResult.Success = success
	}
	return nil
}
func newUserInfoArgs() interface{} {
	return &UserInfoArgs{}
}

func newUserInfoResult() interface{} {
	return &UserInfoResult{}
}

type UserInfoArgs struct {
	Req *user.UserInfoRequest
}

func (p *UserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UserInfoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UserInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UserInfoArgs) Unmarshal(in []byte) error {
	msg := new(user.UserInfoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserInfoArgs_Req_DEFAULT *user.UserInfoRequest

func (p *UserInfoArgs) GetReq() *user.UserInfoRequest {
	if !p.IsSetReq() {
		return UserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

type UserInfoResult struct {
	Success *user.UserInfoResponse
}

var UserInfoResult_Success_DEFAULT *user.UserInfoResponse

func (p *UserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UserInfoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UserInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UserInfoResult) Unmarshal(in []byte) error {
	msg := new(user.UserInfoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserInfoResult) GetSuccess() *user.UserInfoResponse {
	if !p.IsSetSuccess() {
		return UserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UserInfoResponse)
}

func (p *UserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func userLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UserLoginRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UserLogin(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserLoginArgs:
		success, err := handler.(user.UserService).UserLogin(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserLoginResult)
		realResult.Success = success
	}
	return nil
}
func newUserLoginArgs() interface{} {
	return &UserLoginArgs{}
}

func newUserLoginResult() interface{} {
	return &UserLoginResult{}
}

type UserLoginArgs struct {
	Req *user.UserLoginRequest
}

func (p *UserLoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UserLoginRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UserLoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UserLoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UserLoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UserLoginArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UserLoginArgs) Unmarshal(in []byte) error {
	msg := new(user.UserLoginRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserLoginArgs_Req_DEFAULT *user.UserLoginRequest

func (p *UserLoginArgs) GetReq() *user.UserLoginRequest {
	if !p.IsSetReq() {
		return UserLoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserLoginArgs) IsSetReq() bool {
	return p.Req != nil
}

type UserLoginResult struct {
	Success *user.UserLoginResponse
}

var UserLoginResult_Success_DEFAULT *user.UserLoginResponse

func (p *UserLoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UserLoginResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UserLoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UserLoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UserLoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UserLoginResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UserLoginResult) Unmarshal(in []byte) error {
	msg := new(user.UserLoginResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserLoginResult) GetSuccess() *user.UserLoginResponse {
	if !p.IsSetSuccess() {
		return UserLoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserLoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UserLoginResponse)
}

func (p *UserLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func userRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UserRegisterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UserRegister(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserRegisterArgs:
		success, err := handler.(user.UserService).UserRegister(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserRegisterResult)
		realResult.Success = success
	}
	return nil
}
func newUserRegisterArgs() interface{} {
	return &UserRegisterArgs{}
}

func newUserRegisterResult() interface{} {
	return &UserRegisterResult{}
}

type UserRegisterArgs struct {
	Req *user.UserRegisterRequest
}

func (p *UserRegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UserRegisterRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UserRegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UserRegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UserRegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UserRegisterArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UserRegisterArgs) Unmarshal(in []byte) error {
	msg := new(user.UserRegisterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserRegisterArgs_Req_DEFAULT *user.UserRegisterRequest

func (p *UserRegisterArgs) GetReq() *user.UserRegisterRequest {
	if !p.IsSetReq() {
		return UserRegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserRegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

type UserRegisterResult struct {
	Success *user.UserRegisterResponse
}

var UserRegisterResult_Success_DEFAULT *user.UserRegisterResponse

func (p *UserRegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UserRegisterResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UserRegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UserRegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UserRegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UserRegisterResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UserRegisterResult) Unmarshal(in []byte) error {
	msg := new(user.UserRegisterResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserRegisterResult) GetSuccess() *user.UserRegisterResponse {
	if !p.IsSetSuccess() {
		return UserRegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserRegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UserRegisterResponse)
}

func (p *UserRegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getAvatarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UserAvatarRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetAvatar(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetAvatarArgs:
		success, err := handler.(user.UserService).GetAvatar(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetAvatarResult)
		realResult.Success = success
	}
	return nil
}
func newGetAvatarArgs() interface{} {
	return &GetAvatarArgs{}
}

func newGetAvatarResult() interface{} {
	return &GetAvatarResult{}
}

type GetAvatarArgs struct {
	Req *user.UserAvatarRequest
}

func (p *GetAvatarArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UserAvatarRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetAvatarArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetAvatarArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetAvatarArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetAvatarArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetAvatarArgs) Unmarshal(in []byte) error {
	msg := new(user.UserAvatarRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetAvatarArgs_Req_DEFAULT *user.UserAvatarRequest

func (p *GetAvatarArgs) GetReq() *user.UserAvatarRequest {
	if !p.IsSetReq() {
		return GetAvatarArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAvatarArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetAvatarResult struct {
	Success *user.UserAvatarResponse
}

var GetAvatarResult_Success_DEFAULT *user.UserAvatarResponse

func (p *GetAvatarResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UserAvatarResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetAvatarResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetAvatarResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetAvatarResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetAvatarResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetAvatarResult) Unmarshal(in []byte) error {
	msg := new(user.UserAvatarResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAvatarResult) GetSuccess() *user.UserAvatarResponse {
	if !p.IsSetSuccess() {
		return GetAvatarResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAvatarResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UserAvatarResponse)
}

func (p *GetAvatarResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getBackgroundImageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UserBackgroundImageRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetBackgroundImage(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetBackgroundImageArgs:
		success, err := handler.(user.UserService).GetBackgroundImage(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetBackgroundImageResult)
		realResult.Success = success
	}
	return nil
}
func newGetBackgroundImageArgs() interface{} {
	return &GetBackgroundImageArgs{}
}

func newGetBackgroundImageResult() interface{} {
	return &GetBackgroundImageResult{}
}

type GetBackgroundImageArgs struct {
	Req *user.UserBackgroundImageRequest
}

func (p *GetBackgroundImageArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UserBackgroundImageRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetBackgroundImageArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetBackgroundImageArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetBackgroundImageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetBackgroundImageArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetBackgroundImageArgs) Unmarshal(in []byte) error {
	msg := new(user.UserBackgroundImageRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetBackgroundImageArgs_Req_DEFAULT *user.UserBackgroundImageRequest

func (p *GetBackgroundImageArgs) GetReq() *user.UserBackgroundImageRequest {
	if !p.IsSetReq() {
		return GetBackgroundImageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetBackgroundImageArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetBackgroundImageResult struct {
	Success *user.UserBackgroundImageResponse
}

var GetBackgroundImageResult_Success_DEFAULT *user.UserBackgroundImageResponse

func (p *GetBackgroundImageResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UserBackgroundImageResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetBackgroundImageResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetBackgroundImageResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetBackgroundImageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetBackgroundImageResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetBackgroundImageResult) Unmarshal(in []byte) error {
	msg := new(user.UserBackgroundImageResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetBackgroundImageResult) GetSuccess() *user.UserBackgroundImageResponse {
	if !p.IsSetSuccess() {
		return GetBackgroundImageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetBackgroundImageResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UserBackgroundImageResponse)
}

func (p *GetBackgroundImageResult) IsSetSuccess() bool {
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

func (p *kClient) UserInfo(ctx context.Context, Req *user.UserInfoRequest) (r *user.UserInfoResponse, err error) {
	var _args UserInfoArgs
	_args.Req = Req
	var _result UserInfoResult
	if err = p.c.Call(ctx, "UserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserLogin(ctx context.Context, Req *user.UserLoginRequest) (r *user.UserLoginResponse, err error) {
	var _args UserLoginArgs
	_args.Req = Req
	var _result UserLoginResult
	if err = p.c.Call(ctx, "UserLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserRegister(ctx context.Context, Req *user.UserRegisterRequest) (r *user.UserRegisterResponse, err error) {
	var _args UserRegisterArgs
	_args.Req = Req
	var _result UserRegisterResult
	if err = p.c.Call(ctx, "UserRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAvatar(ctx context.Context, Req *user.UserAvatarRequest) (r *user.UserAvatarResponse, err error) {
	var _args GetAvatarArgs
	_args.Req = Req
	var _result GetAvatarResult
	if err = p.c.Call(ctx, "GetAvatar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetBackgroundImage(ctx context.Context, Req *user.UserBackgroundImageRequest) (r *user.UserBackgroundImageResponse, err error) {
	var _args GetBackgroundImageArgs
	_args.Req = Req
	var _result GetBackgroundImageResult
	if err = p.c.Call(ctx, "GetBackgroundImage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
