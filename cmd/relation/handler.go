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
	err = errors.New("未完成")
	return
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	userId := req.UserId
	var myId int64 = 1 // TODO 用户微服务的方法
	UserList, err := service.GetFollowList(ctx, userId, myId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		return
	}
	resp = new(relation.RelationFollowListResponse)
	for _, u := range *UserList {
		resp.UserList = append(resp.UserList, &u)
	}
	resp.StatusCode = 0
	resp.StatusMsg = "ok"
	return
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	err = errors.New("未完成")
	return
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	err = errors.New("未完成")
	return
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	err = errors.New("未完成")
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
