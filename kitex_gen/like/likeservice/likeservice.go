// Code generated by Kitex v0.4.4. DO NOT EDIT.

package likeservice

import (
	"context"
	like "dousheng/kitex_gen/like"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return likeServiceServiceInfo
}

var likeServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "LikeService"
	handlerType := (*like.LikeService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteAction":   kitex.NewMethodInfo(favoriteActionHandler, newFavoriteActionArgs, newFavoriteActionResult, false),
		"GetFavoriteList":  kitex.NewMethodInfo(getFavoriteListHandler, newGetFavoriteListArgs, newGetFavoriteListResult, false),
		"TotalFavorite":    kitex.NewMethodInfo(totalFavoriteHandler, newTotalFavoriteArgs, newTotalFavoriteResult, false),
		"FavoriteCount":    kitex.NewMethodInfo(favoriteCountHandler, newFavoriteCountArgs, newFavoriteCountResult, false),
		"isFavorite":       kitex.NewMethodInfo(isFavoriteHandler, newIsFavoriteArgs, newIsFavoriteResult, false),
		"RefreshLikeCache": kitex.NewMethodInfo(refreshLikeCacheHandler, newRefreshLikeCacheArgs, newRefreshLikeCacheResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "like",
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

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(like.FavoriteActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(like.LikeService).FavoriteAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteActionArgs:
		success, err := handler.(like.LikeService).FavoriteAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteActionResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteActionArgs() interface{} {
	return &FavoriteActionArgs{}
}

func newFavoriteActionResult() interface{} {
	return &FavoriteActionResult{}
}

type FavoriteActionArgs struct {
	Req *like.FavoriteActionRequest
}

func (p *FavoriteActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(like.FavoriteActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoriteActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoriteActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoriteActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteActionArgs) Unmarshal(in []byte) error {
	msg := new(like.FavoriteActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteActionArgs_Req_DEFAULT *like.FavoriteActionRequest

func (p *FavoriteActionArgs) GetReq() *like.FavoriteActionRequest {
	if !p.IsSetReq() {
		return FavoriteActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type FavoriteActionResult struct {
	Success *like.FavoriteActionResponse
}

var FavoriteActionResult_Success_DEFAULT *like.FavoriteActionResponse

func (p *FavoriteActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(like.FavoriteActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoriteActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoriteActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoriteActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteActionResult) Unmarshal(in []byte) error {
	msg := new(like.FavoriteActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteActionResult) GetSuccess() *like.FavoriteActionResponse {
	if !p.IsSetSuccess() {
		return FavoriteActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*like.FavoriteActionResponse)
}

func (p *FavoriteActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getFavoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(like.FavoriteListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(like.LikeService).GetFavoriteList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetFavoriteListArgs:
		success, err := handler.(like.LikeService).GetFavoriteList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFavoriteListResult)
		realResult.Success = success
	}
	return nil
}
func newGetFavoriteListArgs() interface{} {
	return &GetFavoriteListArgs{}
}

func newGetFavoriteListResult() interface{} {
	return &GetFavoriteListResult{}
}

type GetFavoriteListArgs struct {
	Req *like.FavoriteListRequest
}

func (p *GetFavoriteListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(like.FavoriteListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetFavoriteListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetFavoriteListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetFavoriteListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFavoriteListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetFavoriteListArgs) Unmarshal(in []byte) error {
	msg := new(like.FavoriteListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFavoriteListArgs_Req_DEFAULT *like.FavoriteListRequest

func (p *GetFavoriteListArgs) GetReq() *like.FavoriteListRequest {
	if !p.IsSetReq() {
		return GetFavoriteListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFavoriteListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetFavoriteListResult struct {
	Success *like.FavoriteListResponse
}

var GetFavoriteListResult_Success_DEFAULT *like.FavoriteListResponse

func (p *GetFavoriteListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(like.FavoriteListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetFavoriteListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetFavoriteListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetFavoriteListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFavoriteListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetFavoriteListResult) Unmarshal(in []byte) error {
	msg := new(like.FavoriteListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFavoriteListResult) GetSuccess() *like.FavoriteListResponse {
	if !p.IsSetSuccess() {
		return GetFavoriteListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFavoriteListResult) SetSuccess(x interface{}) {
	p.Success = x.(*like.FavoriteListResponse)
}

func (p *GetFavoriteListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func totalFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(like.TotalFavoriteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(like.LikeService).TotalFavorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *TotalFavoriteArgs:
		success, err := handler.(like.LikeService).TotalFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*TotalFavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newTotalFavoriteArgs() interface{} {
	return &TotalFavoriteArgs{}
}

func newTotalFavoriteResult() interface{} {
	return &TotalFavoriteResult{}
}

type TotalFavoriteArgs struct {
	Req *like.TotalFavoriteRequest
}

func (p *TotalFavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(like.TotalFavoriteRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *TotalFavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *TotalFavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *TotalFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in TotalFavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *TotalFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(like.TotalFavoriteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var TotalFavoriteArgs_Req_DEFAULT *like.TotalFavoriteRequest

func (p *TotalFavoriteArgs) GetReq() *like.TotalFavoriteRequest {
	if !p.IsSetReq() {
		return TotalFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *TotalFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

type TotalFavoriteResult struct {
	Success *like.TotalFavoriteResponse
}

var TotalFavoriteResult_Success_DEFAULT *like.TotalFavoriteResponse

func (p *TotalFavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(like.TotalFavoriteResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *TotalFavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *TotalFavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *TotalFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in TotalFavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *TotalFavoriteResult) Unmarshal(in []byte) error {
	msg := new(like.TotalFavoriteResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *TotalFavoriteResult) GetSuccess() *like.TotalFavoriteResponse {
	if !p.IsSetSuccess() {
		return TotalFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *TotalFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*like.TotalFavoriteResponse)
}

func (p *TotalFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func favoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(like.FavoriteCountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(like.LikeService).FavoriteCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteCountArgs:
		success, err := handler.(like.LikeService).FavoriteCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteCountResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteCountArgs() interface{} {
	return &FavoriteCountArgs{}
}

func newFavoriteCountResult() interface{} {
	return &FavoriteCountResult{}
}

type FavoriteCountArgs struct {
	Req *like.FavoriteCountRequest
}

func (p *FavoriteCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(like.FavoriteCountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoriteCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoriteCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoriteCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteCountArgs) Unmarshal(in []byte) error {
	msg := new(like.FavoriteCountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteCountArgs_Req_DEFAULT *like.FavoriteCountRequest

func (p *FavoriteCountArgs) GetReq() *like.FavoriteCountRequest {
	if !p.IsSetReq() {
		return FavoriteCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteCountArgs) IsSetReq() bool {
	return p.Req != nil
}

type FavoriteCountResult struct {
	Success *like.FavoriteCountResponse
}

var FavoriteCountResult_Success_DEFAULT *like.FavoriteCountResponse

func (p *FavoriteCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(like.FavoriteCountResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoriteCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoriteCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoriteCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteCountResult) Unmarshal(in []byte) error {
	msg := new(like.FavoriteCountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteCountResult) GetSuccess() *like.FavoriteCountResponse {
	if !p.IsSetSuccess() {
		return FavoriteCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*like.FavoriteCountResponse)
}

func (p *FavoriteCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func isFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(like.IsFavoriteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(like.LikeService).IsFavorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *IsFavoriteArgs:
		success, err := handler.(like.LikeService).IsFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*IsFavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newIsFavoriteArgs() interface{} {
	return &IsFavoriteArgs{}
}

func newIsFavoriteResult() interface{} {
	return &IsFavoriteResult{}
}

type IsFavoriteArgs struct {
	Req *like.IsFavoriteRequest
}

func (p *IsFavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(like.IsFavoriteRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *IsFavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *IsFavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *IsFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in IsFavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *IsFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(like.IsFavoriteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var IsFavoriteArgs_Req_DEFAULT *like.IsFavoriteRequest

func (p *IsFavoriteArgs) GetReq() *like.IsFavoriteRequest {
	if !p.IsSetReq() {
		return IsFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *IsFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

type IsFavoriteResult struct {
	Success *like.IsFavoriteResponse
}

var IsFavoriteResult_Success_DEFAULT *like.IsFavoriteResponse

func (p *IsFavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(like.IsFavoriteResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *IsFavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *IsFavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *IsFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in IsFavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *IsFavoriteResult) Unmarshal(in []byte) error {
	msg := new(like.IsFavoriteResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *IsFavoriteResult) GetSuccess() *like.IsFavoriteResponse {
	if !p.IsSetSuccess() {
		return IsFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *IsFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*like.IsFavoriteResponse)
}

func (p *IsFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func refreshLikeCacheHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(like.RefreshLikeCacheRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(like.LikeService).RefreshLikeCache(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RefreshLikeCacheArgs:
		success, err := handler.(like.LikeService).RefreshLikeCache(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RefreshLikeCacheResult)
		realResult.Success = success
	}
	return nil
}
func newRefreshLikeCacheArgs() interface{} {
	return &RefreshLikeCacheArgs{}
}

func newRefreshLikeCacheResult() interface{} {
	return &RefreshLikeCacheResult{}
}

type RefreshLikeCacheArgs struct {
	Req *like.RefreshLikeCacheRequest
}

func (p *RefreshLikeCacheArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(like.RefreshLikeCacheRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RefreshLikeCacheArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RefreshLikeCacheArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RefreshLikeCacheArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in RefreshLikeCacheArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *RefreshLikeCacheArgs) Unmarshal(in []byte) error {
	msg := new(like.RefreshLikeCacheRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RefreshLikeCacheArgs_Req_DEFAULT *like.RefreshLikeCacheRequest

func (p *RefreshLikeCacheArgs) GetReq() *like.RefreshLikeCacheRequest {
	if !p.IsSetReq() {
		return RefreshLikeCacheArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RefreshLikeCacheArgs) IsSetReq() bool {
	return p.Req != nil
}

type RefreshLikeCacheResult struct {
	Success *like.RefreshLikeCacheResponse
}

var RefreshLikeCacheResult_Success_DEFAULT *like.RefreshLikeCacheResponse

func (p *RefreshLikeCacheResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(like.RefreshLikeCacheResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RefreshLikeCacheResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RefreshLikeCacheResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RefreshLikeCacheResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in RefreshLikeCacheResult")
	}
	return proto.Marshal(p.Success)
}

func (p *RefreshLikeCacheResult) Unmarshal(in []byte) error {
	msg := new(like.RefreshLikeCacheResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RefreshLikeCacheResult) GetSuccess() *like.RefreshLikeCacheResponse {
	if !p.IsSetSuccess() {
		return RefreshLikeCacheResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RefreshLikeCacheResult) SetSuccess(x interface{}) {
	p.Success = x.(*like.RefreshLikeCacheResponse)
}

func (p *RefreshLikeCacheResult) IsSetSuccess() bool {
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

func (p *kClient) FavoriteAction(ctx context.Context, Req *like.FavoriteActionRequest) (r *like.FavoriteActionResponse, err error) {
	var _args FavoriteActionArgs
	_args.Req = Req
	var _result FavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFavoriteList(ctx context.Context, Req *like.FavoriteListRequest) (r *like.FavoriteListResponse, err error) {
	var _args GetFavoriteListArgs
	_args.Req = Req
	var _result GetFavoriteListResult
	if err = p.c.Call(ctx, "GetFavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TotalFavorite(ctx context.Context, Req *like.TotalFavoriteRequest) (r *like.TotalFavoriteResponse, err error) {
	var _args TotalFavoriteArgs
	_args.Req = Req
	var _result TotalFavoriteResult
	if err = p.c.Call(ctx, "TotalFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteCount(ctx context.Context, Req *like.FavoriteCountRequest) (r *like.FavoriteCountResponse, err error) {
	var _args FavoriteCountArgs
	_args.Req = Req
	var _result FavoriteCountResult
	if err = p.c.Call(ctx, "FavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsFavorite(ctx context.Context, Req *like.IsFavoriteRequest) (r *like.IsFavoriteResponse, err error) {
	var _args IsFavoriteArgs
	_args.Req = Req
	var _result IsFavoriteResult
	if err = p.c.Call(ctx, "isFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RefreshLikeCache(ctx context.Context, Req *like.RefreshLikeCacheRequest) (r *like.RefreshLikeCacheResponse, err error) {
	var _args RefreshLikeCacheArgs
	_args.Req = Req
	var _result RefreshLikeCacheResult
	if err = p.c.Call(ctx, "RefreshLikeCache", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
