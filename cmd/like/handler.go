package main

import (
	"context"
	like "dousheng/kitex_gen/like"
	"errors"
)

// LikeServiceImpl implements the last service interface defined in the IDL.
type LikeServiceImpl struct{}

// FavoriteAction implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) FavoriteAction(ctx context.Context, req *like.FavoriteActionRequest) (resp *like.FavoriteActionResponse, err error) {
	err = errors.New("未完成")
	return
}

// GetFavoriteList implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) GetFavoriteList(ctx context.Context, req *like.FavoriteListRequest) (resp *like.FavoriteListResponse, err error) {
	err = errors.New("未完成")
	return
}

// FavoriteCount implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) FavoriteCount(ctx context.Context, req *like.FavoriteCountRequest) (resp *like.FavoriteCountResponse, err error) {
	err = errors.New("未完成")
	return
}
