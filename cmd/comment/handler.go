package main

import (
	"context"
	comment "dousheng/kitex_gen/comment"
	"errors"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// PostCommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) PostCommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	err = errors.New("未完成")
	return
}

// GetCommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetCommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	err = errors.New("未完成")
	return
}
