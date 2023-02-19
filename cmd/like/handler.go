package main

import (
	"context"
	service "dousheng/cmd/like/internal/service"
	like "dousheng/kitex_gen/like"
)

// LikeServiceImpl implements the last service interface defined in the IDL.
type LikeServiceImpl struct{}

// FavoriteAction implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) FavoriteAction(ctx context.Context, req *like.FavoriteActionRequest) (resp *like.FavoriteActionResponse, err error) {
	err = service.FavoriteAction(req.GetUserId(), req.GetVideoId(), req.GetActionType())
	if err != nil {
		resp.StatusCode = 0
		resp.StatusMsg = "ok"
	}
	return
}

// GetFavoriteList implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) GetFavoriteList(ctx context.Context, req *like.FavoriteListRequest) (resp *like.FavoriteListResponse, err error) {
	resp = &like.FavoriteListResponse{}
	resp.VideoList, err = service.GetFavoriteList(ctx, req.GetUserId())
	if err != nil {
		resp.StatusCode = 0
		resp.StatusMsg = "ok"
	}
	return
}

// TotalFavorite implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) TotalFavorite(ctx context.Context, req *like.TotalFavoriteRequest) (resp *like.TotalFavoriteResponse, err error) {
	resp = &like.TotalFavoriteResponse{}
	fWorks := service.FavoriteVideoCount(req.UserId)
	tfd := service.TotalFavoriteCount(req.UserId)
	resp.TotalFavorited = tfd
	resp.FavoriteCount = fWorks
	return
}

// FavoriteCount implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) FavoriteCount(ctx context.Context, req *like.FavoriteCountRequest) (resp *like.FavoriteCountResponse, err error) {
	resp = &like.FavoriteCountResponse{}
	vfc := service.VideoFavoriteCount(req.VideoId)
	resp.FavoriteCount = vfc
	return
}

// IsFavorite implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) IsFavorite(ctx context.Context, req *like.IsFavoriteRequest) (resp *like.IsFavoriteResponse, err error) {
	resp = &like.IsFavoriteResponse{}
	isf := service.IsLike(req.VideoId, req.UserId)
	resp.IsFavorite = isf
	return
}
