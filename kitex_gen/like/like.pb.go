// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: like.proto

package like

import (
	context "context"
	video "dousheng/kitex_gen/video"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 喜欢操作
type FavoriteActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`             // 用户id
	VideoId    int64 `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`          // 视频id
	ActionType int32 `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"` // 1-点赞，2-取消点赞
}

func (x *FavoriteActionRequest) Reset() {
	*x = FavoriteActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteActionRequest) ProtoMessage() {}

func (x *FavoriteActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteActionRequest.ProtoReflect.Descriptor instead.
func (*FavoriteActionRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{0}
}

func (x *FavoriteActionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FavoriteActionRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *FavoriteActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type FavoriteActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
}

func (x *FavoriteActionResponse) Reset() {
	*x = FavoriteActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteActionResponse) ProtoMessage() {}

func (x *FavoriteActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteActionResponse.ProtoReflect.Descriptor instead.
func (*FavoriteActionResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{1}
}

func (x *FavoriteActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FavoriteActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

// 喜欢列表
type FavoriteListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
	MyId   int64 `protobuf:"varint,2,opt,name=my_id,json=myId,proto3" json:"my_id,omitempty"`       // 登录的用户id
}

func (x *FavoriteListRequest) Reset() {
	*x = FavoriteListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteListRequest) ProtoMessage() {}

func (x *FavoriteListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteListRequest.ProtoReflect.Descriptor instead.
func (*FavoriteListRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{2}
}

func (x *FavoriteListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FavoriteListRequest) GetMyId() int64 {
	if x != nil {
		return x.MyId
	}
	return 0
}

type FavoriteListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32          `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string         `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	VideoList  []*video.Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`     // 用户点赞视频列表
}

func (x *FavoriteListResponse) Reset() {
	*x = FavoriteListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteListResponse) ProtoMessage() {}

func (x *FavoriteListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteListResponse.ProtoReflect.Descriptor instead.
func (*FavoriteListResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{3}
}

func (x *FavoriteListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FavoriteListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *FavoriteListResponse) GetVideoList() []*video.Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

// 喜欢总数
type TotalFavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
}

func (x *TotalFavoriteRequest) Reset() {
	*x = TotalFavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalFavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalFavoriteRequest) ProtoMessage() {}

func (x *TotalFavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalFavoriteRequest.ProtoReflect.Descriptor instead.
func (*TotalFavoriteRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{4}
}

func (x *TotalFavoriteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type TotalFavoriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalFavorited int32 `protobuf:"varint,1,opt,name=total_favorited,json=totalFavorited,proto3" json:"total_favorited,omitempty"` //获赞总数量
	FavoriteCount  int32 `protobuf:"varint,2,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`    //点赞作品总数量
}

func (x *TotalFavoriteResponse) Reset() {
	*x = TotalFavoriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalFavoriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalFavoriteResponse) ProtoMessage() {}

func (x *TotalFavoriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalFavoriteResponse.ProtoReflect.Descriptor instead.
func (*TotalFavoriteResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{5}
}

func (x *TotalFavoriteResponse) GetTotalFavorited() int32 {
	if x != nil {
		return x.TotalFavorited
	}
	return 0
}

func (x *TotalFavoriteResponse) GetFavoriteCount() int32 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

// 视频点赞数
type FavoriteCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int64 `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"` // 视频id
}

func (x *FavoriteCountRequest) Reset() {
	*x = FavoriteCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteCountRequest) ProtoMessage() {}

func (x *FavoriteCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteCountRequest.ProtoReflect.Descriptor instead.
func (*FavoriteCountRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{6}
}

func (x *FavoriteCountRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type FavoriteCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FavoriteCount int32 `protobuf:"varint,1,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"` //视频的点赞数量
}

func (x *FavoriteCountResponse) Reset() {
	*x = FavoriteCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteCountResponse) ProtoMessage() {}

func (x *FavoriteCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteCountResponse.ProtoReflect.Descriptor instead.
func (*FavoriteCountResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{7}
}

func (x *FavoriteCountResponse) GetFavoriteCount() int32 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

// 视频是否点赞
type IsFavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`    // 用户id
	VideoId int64 `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"` // 视频id
}

func (x *IsFavoriteRequest) Reset() {
	*x = IsFavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsFavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsFavoriteRequest) ProtoMessage() {}

func (x *IsFavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsFavoriteRequest.ProtoReflect.Descriptor instead.
func (*IsFavoriteRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{8}
}

func (x *IsFavoriteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *IsFavoriteRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type IsFavoriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFavorite bool `protobuf:"varint,1,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"` // true-已点赞，false-未点赞
}

func (x *IsFavoriteResponse) Reset() {
	*x = IsFavoriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsFavoriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsFavoriteResponse) ProtoMessage() {}

func (x *IsFavoriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsFavoriteResponse.ProtoReflect.Descriptor instead.
func (*IsFavoriteResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{9}
}

func (x *IsFavoriteResponse) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

type RefreshLikeCacheRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshLikeCacheRequest) Reset() {
	*x = RefreshLikeCacheRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshLikeCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshLikeCacheRequest) ProtoMessage() {}

func (x *RefreshLikeCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshLikeCacheRequest.ProtoReflect.Descriptor instead.
func (*RefreshLikeCacheRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{10}
}

type RefreshLikeCacheResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshLikeCacheResponse) Reset() {
	*x = RefreshLikeCacheResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshLikeCacheResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshLikeCacheResponse) ProtoMessage() {}

func (x *RefreshLikeCacheResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshLikeCacheResponse.ProtoReflect.Descriptor instead.
func (*RefreshLikeCacheResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{11}
}

var File_like_proto protoreflect.FileDescriptor

var file_like_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x69,
	0x6b, 0x65, 0x1a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6c, 0x0a, 0x15, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x58, 0x0a,
	0x16, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x43, 0x0a, 0x13, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x6d, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6d, 0x79, 0x49, 0x64, 0x22, 0x83, 0x01, 0x0a,
	0x14, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x2b, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x2f, 0x0a, 0x14, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x15, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x31, 0x0a, 0x14,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22,
	0x3e, 0x0a, 0x15, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x47, 0x0a, 0x11, 0x49, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x12, 0x49, 0x73, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x22,
	0x19, 0x0a, 0x17, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd8, 0x03, 0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4a, 0x0a, 0x0d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x12, 0x1a, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a,
	0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x69, 0x6b,
	0x65, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x49, 0x73, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x49,
	0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x49, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x10,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x12, 0x1d, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4c,
	0x69, 0x6b, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4c, 0x69,
	0x6b, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x19, 0x5a, 0x17, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x6b, 0x69,
	0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_like_proto_rawDescOnce sync.Once
	file_like_proto_rawDescData = file_like_proto_rawDesc
)

func file_like_proto_rawDescGZIP() []byte {
	file_like_proto_rawDescOnce.Do(func() {
		file_like_proto_rawDescData = protoimpl.X.CompressGZIP(file_like_proto_rawDescData)
	})
	return file_like_proto_rawDescData
}

var file_like_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_like_proto_goTypes = []interface{}{
	(*FavoriteActionRequest)(nil),    // 0: like.FavoriteActionRequest
	(*FavoriteActionResponse)(nil),   // 1: like.FavoriteActionResponse
	(*FavoriteListRequest)(nil),      // 2: like.FavoriteListRequest
	(*FavoriteListResponse)(nil),     // 3: like.FavoriteListResponse
	(*TotalFavoriteRequest)(nil),     // 4: like.TotalFavoriteRequest
	(*TotalFavoriteResponse)(nil),    // 5: like.TotalFavoriteResponse
	(*FavoriteCountRequest)(nil),     // 6: like.FavoriteCountRequest
	(*FavoriteCountResponse)(nil),    // 7: like.FavoriteCountResponse
	(*IsFavoriteRequest)(nil),        // 8: like.IsFavoriteRequest
	(*IsFavoriteResponse)(nil),       // 9: like.IsFavoriteResponse
	(*RefreshLikeCacheRequest)(nil),  // 10: like.RefreshLikeCacheRequest
	(*RefreshLikeCacheResponse)(nil), // 11: like.RefreshLikeCacheResponse
	(*video.Video)(nil),              // 12: video.Video
}
var file_like_proto_depIdxs = []int32{
	12, // 0: like.FavoriteListResponse.video_list:type_name -> video.Video
	0,  // 1: like.LikeService.FavoriteAction:input_type -> like.FavoriteActionRequest
	2,  // 2: like.LikeService.GetFavoriteList:input_type -> like.FavoriteListRequest
	4,  // 3: like.LikeService.TotalFavorite:input_type -> like.TotalFavoriteRequest
	6,  // 4: like.LikeService.FavoriteCount:input_type -> like.FavoriteCountRequest
	8,  // 5: like.LikeService.IsFavorite:input_type -> like.IsFavoriteRequest
	10, // 6: like.LikeService.RefreshLikeCache:input_type -> like.RefreshLikeCacheRequest
	1,  // 7: like.LikeService.FavoriteAction:output_type -> like.FavoriteActionResponse
	3,  // 8: like.LikeService.GetFavoriteList:output_type -> like.FavoriteListResponse
	5,  // 9: like.LikeService.TotalFavorite:output_type -> like.TotalFavoriteResponse
	7,  // 10: like.LikeService.FavoriteCount:output_type -> like.FavoriteCountResponse
	9,  // 11: like.LikeService.IsFavorite:output_type -> like.IsFavoriteResponse
	11, // 12: like.LikeService.RefreshLikeCache:output_type -> like.RefreshLikeCacheResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_like_proto_init() }
func file_like_proto_init() {
	if File_like_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_like_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteActionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalFavoriteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalFavoriteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteCountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteCountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsFavoriteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsFavoriteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshLikeCacheRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshLikeCacheResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_like_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_like_proto_goTypes,
		DependencyIndexes: file_like_proto_depIdxs,
		MessageInfos:      file_like_proto_msgTypes,
	}.Build()
	File_like_proto = out.File
	file_like_proto_rawDesc = nil
	file_like_proto_goTypes = nil
	file_like_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.4.4. DO NOT EDIT.

type LikeService interface {
	FavoriteAction(ctx context.Context, req *FavoriteActionRequest) (res *FavoriteActionResponse, err error)
	GetFavoriteList(ctx context.Context, req *FavoriteListRequest) (res *FavoriteListResponse, err error)
	TotalFavorite(ctx context.Context, req *TotalFavoriteRequest) (res *TotalFavoriteResponse, err error)
	FavoriteCount(ctx context.Context, req *FavoriteCountRequest) (res *FavoriteCountResponse, err error)
	IsFavorite(ctx context.Context, req *IsFavoriteRequest) (res *IsFavoriteResponse, err error)
	RefreshLikeCache(ctx context.Context, req *RefreshLikeCacheRequest) (res *RefreshLikeCacheResponse, err error)
}
