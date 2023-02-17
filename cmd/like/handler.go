package main

import (
	"context"
	like "dousheng/kitex_gen/like"
)

// LikeServiceImpl implements the last service interface defined in the IDL.
type LikeServiceImpl struct{}

// FavoriteAction implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) FavoriteAction(ctx context.Context, req *like.FavoriteActionRequest) (resp *like.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFavoriteList implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) GetFavoriteList(ctx context.Context, req *like.FavoriteListRequest) (resp *like.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}
