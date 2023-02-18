package main

import (
	"context"
	"dousheng/cmd/relation/internal/service"
	relation "dousheng/kitex_gen/relation"
	"github.com/pkg/errors"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	err = service.RelationAction(req.GetFromUserId(), req.GetToUserId(), req.GetActionType())
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
	}
	resp.StatusCode = 0
	resp.StatusMsg = "ok"
	return
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	resp, err = service.GetFollowList(ctx, req.GetUserId(), req.GetMyId())
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
	}
	resp.StatusCode = 0
	resp.StatusMsg = "ok"
	return
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	resp, err = service.GetFollowerList(ctx, req.GetUserId(), req.GetMyId())
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
	}
	resp.StatusCode = 0
	resp.StatusMsg = "ok"
	return
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	resp, err = service.GetFriendList(ctx, req.GetUserId(), req.GetMyId())
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
	}
	resp.StatusCode = 0
	resp.StatusMsg = "ok"
	return
}

// GetFollowCount implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowCount(ctx context.Context, req *relation.RelationFollowCountRequest) (resp *relation.RelationFollowCountResponse, err error) {
	err = errors.New("未完成")
	return
}

// GetFollowerCount implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerCount(ctx context.Context, req *relation.RelationFollowerCountRequest) (resp *relation.RelationFollowerCountResponse, err error) {
	err = errors.New("未完成")
	return
}

// GetIsFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetIsFollow(ctx context.Context, req *relation.RelationIsFollowRequest) (resp *relation.RelationIsFollowResponse, err error) {

	return
}

// GetFriendsMessageList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendsMessageList(ctx context.Context, req *relation.RelationFriendsMessageListRequest) (resp *relation.RelationFriendsMessageListResponse, err error) {
	// TODO: Your code here...
	return
}
