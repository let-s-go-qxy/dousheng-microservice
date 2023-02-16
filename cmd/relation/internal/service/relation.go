package service

import (
	"context"
	"dousheng/kitex_gen/user"
	"dousheng/pkg/etcd_discovery"
)

func GetFollowList(ctx context.Context, userId int64, myId int64) (*[]user.User, error) {
	resp, err := etcd_discovery.UserInfo(ctx, &user.UserInfoRequest{
		UserId: userId,
		Token:  "xxx",
	})
	var followUsers []user.User
	followUsers = append(followUsers, user.User{
		Id:            resp.User.Id,
		Name:          resp.User.Name,
		FollowCount:   resp.User.FollowCount,
		FollowerCount: resp.User.FollowerCount,
		IsFollow:      resp.User.IsFollow,
	})
	return &followUsers, err
}
