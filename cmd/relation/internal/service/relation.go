package service

import (
	"context"
	"dousheng/cmd/relation/internal/model"
	"dousheng/kitex_gen/message"
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

func GetFollowCount(ctx context.Context, userId int64) (*relation.RelationFollowCountResponse, error) {
	count := model.GetFollowCount(userId)
	return &relation.RelationFollowCountResponse{
		Count: int32(count),
	}, nil
}

func GetFollowerCount(ctx context.Context, userId int64) (*relation.RelationFollowerCountResponse, error) {
	count := model.GetFollowerCount(userId)
	return &relation.RelationFollowerCountResponse{
		Count: int32(count),
	}, nil
}

func IsFollow(ctx context.Context, userId int64, myId int64) (*relation.RelationIsFollowResponse, error) {
	return &relation.RelationIsFollowResponse{
		IsFollow: model.IsFollow(myId, userId),
	}, nil
}

func GetFriendMessageList(ctx context.Context, userId int64) (*relation.RelationFriendsMessageListResponse, error) {
	ids := model.GetFriendsByUserId(userId)
	messageList := &relation.RelationFriendsMessageListResponse{}
	for _, friendId := range ids {
		chat, err := etcd_discovery.MessageClient.GetMessageList(ctx, &message.MessageChatRequest{
			UserId:   userId,
			ToUserId: friendId,
		})
		if err != nil {
			return nil, err
		}
		//chat2, err := etcd_discovery.MessageClient.GetMessageList(ctx, &message.MessageChatRequest{
		//	UserId:   friendId,
		//	ToUserId: userId,
		//})
		for _, msg := range chat.GetMessageList() {
			messageList.MessageList = append(messageList.MessageList, &message.Message{
				Id:         msg.GetId(),
				ToUserId:   msg.GetToUserId(),
				FromUserId: msg.GetFromUserId(),
				Content:    msg.GetContent(),
				CreateTime: msg.GetCreateTime(),
			})
		}
	}
	return messageList, nil
}
