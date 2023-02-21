package main

import (
	"context"
	"dousheng/cmd/comment/internal/service"
	comment "dousheng/kitex_gen/comment"
	"dousheng/kitex_gen/user"
	"dousheng/pkg/etcd_discovery"
	g "dousheng/pkg/global"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// PostCommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) PostCommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	resp = &comment.CommentActionResponse{}
	id, content, createDate, err := service.CommentAction(req.VideoId, req.ActionType, req.CommentText, req.CommentId, req.UserId)
	if err != nil {
		return
	}
	info, err := etcd_discovery.UserClient.UserInfo(ctx, &user.UserInfoRequest{
		UserId: req.UserId,
		MyId:   req.UserId,
	})
	if err != nil {
		return
	}
	resp.Comment = &comment.Comment{
		Id:         id,
		User:       info.GetUser(),
		Content:    content,
		CreateDate: createDate,
	}
	resp.StatusCode = g.StatusOk
	resp.StatusMsg = "ok"
	return
}

// GetCommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetCommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	resp = &comment.CommentListResponse{}
	info := &user.UserInfoResponse{}
	commentList := service.GetCommentList(req.VideoId)
	for _, c := range commentList {
		info, err = etcd_discovery.UserClient.UserInfo(ctx, &user.UserInfoRequest{
			UserId: c.User.Id,
			MyId:   req.UserId,
		})
		if err != nil {
			return
		}
		resp.CommentList = append(resp.GetCommentList(), &comment.Comment{
			Id:         c.Id,
			User:       info.User,
			Content:    c.Content,
			CreateDate: c.CreateDate,
		})
	}
	resp.StatusCode = g.StatusOk
	resp.StatusMsg = "ok"
	return
}

// CommentCount implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentCount(ctx context.Context, req *comment.CommentCountRequest) (resp *comment.CommentCountResponse, err error) {
	count := service.GetCommentCount(req.VideoId)
	return &comment.CommentCountResponse{
		CommentCount: count}, nil
}
