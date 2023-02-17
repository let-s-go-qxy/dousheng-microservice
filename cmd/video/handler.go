package main

import (
	"context"
	video "dousheng/kitex_gen/video"
	"errors"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishVideoCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideoCount(ctx context.Context, req *video.PublishVideoCountRequest) (resp *video.PublishVideoCountResponse, err error) {
	// TODO: Your code here...
	err = errors.New("未完成")
	return
}
