package main

import (
	"context"
	"dousheng/cmd/comment/internal/service"
	comment "dousheng/kitex_gen/comment"
	"dousheng/kitex_gen/user"
	"dousheng/pkg/etcd_discovery"
	g "dousheng/pkg/global"

	"github.com/jinzhu/copier"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// PostCommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) PostCommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {

	comments := comment.Comment{}
	user := user.User{}
	comments, err = service.CommentAction(req.VideoId, req.ActionType, req.CommentText, req.CommentId, req.UserId)

	return
}

// GetCommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetCommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {

	info := &user.UserInfoResponse{}

	info, _ = etcd_discovery.UserClient.UserInfo(ctx, &user.UserInfoRequest{
		UserId: req.UserId,
		MyId:   req.UserId,
	})

	commentList := []comment.Comment{}
	commentList = service.GetCommentList(req.VideoId, req.VideoId)
	copier.Copy(resp.CommentList, commentList)
	resp.StatusCode = g.StatusOk
	resp.StatusMsg = "ok"
	return
}
