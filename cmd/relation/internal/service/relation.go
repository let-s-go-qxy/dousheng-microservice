package service

import (
	"context"
	"dousheng/cmd/relation/internal/model"
	"dousheng/kitex_gen/relation"
	"dousheng/kitex_gen/user"
	"dousheng/pkg/etcd_discovery"
)

func GetFollowList(ctx context.Context, userId int64, myId int64) (*relation.RelationFollowListResponse, error) {
	ids := model.GetFollowsByUserId(userId)
	followUsers := new(relation.RelationFollowListResponse)
	for _, id := range ids {
		resp, err := etcd_discovery.UserClient.UserInfo(ctx, &user.UserInfoRequest{
			UserId: id,
			MyId:   myId,
		})
		if err != nil {
			return nil, err
		}
		followUsers.UserList = append(followUsers.UserList, &user.User{
			Id:            resp.User.Id,
			Name:          resp.User.Name,
			FollowCount:   resp.User.FollowCount,
			FollowerCount: resp.User.FollowerCount,
			IsFollow:      resp.User.IsFollow,
		})
	}
	return followUsers, nil
}

func GetFollowerList(ctx context.Context, userId int64, myId int64) (*relation.RelationFollowerListResponse, error) {
	ids := model.GetFollowersByUserId(userId)
	followerUsers := new(relation.RelationFollowerListResponse)
	for _, id := range ids {
		resp, err := etcd_discovery.UserClient.UserInfo(ctx, &user.UserInfoRequest{
			UserId: id,
			MyId:   myId,
		})
		if err != nil {
			return nil, err
		}
		followerUsers.UserList = append(followerUsers.UserList, &user.User{
			Id:            resp.User.Id,
			Name:          resp.User.Name,
			FollowCount:   resp.User.FollowCount,
			FollowerCount: resp.User.FollowerCount,
			IsFollow:      resp.User.IsFollow,
		})
	}
	return followerUsers, nil
}

func GetFriendList(ctx context.Context, userId int64, myId int64) (*relation.RelationFriendListResponse, error) {
	ids := model.GetFriendsByUserId(userId)
	friends := new(relation.RelationFriendListResponse)
	for _, id := range ids {
		resp, err := etcd_discovery.UserClient.UserInfo(ctx, &user.UserInfoRequest{
			UserId: id,
			MyId:   myId,
		})
		if err != nil {
			return nil, err
		}
		friends.UserList = append(friends.UserList, &relation.FriendUser{
			Id:            resp.User.Id,
			Name:          resp.User.Name,
			FollowCount:   resp.User.FollowCount,
			FollowerCount: resp.User.FollowerCount,
			IsFollow:      resp.User.IsFollow,
			Avatar:        resp.User.Avatar,
			Message:       "最后一条信息", // TODO 等待齐迪增加功能
			MsgType:       0,
		})
	}
	return friends, nil
}

func RelationAction(myId, toUserId int64, actionType int32) error {
	return model.CreateOrUpdateFollow(myId, toUserId, actionType)
}
