package service

import (
	"context"
	"dousheng/cmd/relation/internal/model"
	"dousheng/kitex_gen/message"
	"dousheng/kitex_gen/relation"
	"dousheng/kitex_gen/user"
	"dousheng/pkg/etcd_discovery"
	"github.com/cloudwego/kitex/pkg/klog"
	"sort"
)

func GetFollowList(ctx context.Context, userId int64, myId int64) (*relation.RelationFollowListResponse, error) {
	ids := model.GetFollowsByUserId(userId)
	followUsers := new(relation.RelationFollowListResponse)
	//followUsers.UserList = make([]*user.User, 0)
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
	followerUsers.UserList = make([]*user.User, 0)
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

// GetFriendList 获取朋友列表并附带最新的消息
func GetFriendList(ctx context.Context, userId int64, myId int64) (*relation.RelationFriendListResponse, error) {
	ids := model.GetFriendsByUserId(myId)
	friends := new(relation.RelationFriendListResponse)
	friends.UserList = make([]*relation.FriendUser, 0)
	for _, id := range ids {
		//得到myID对应的所有朋友的用户信息并封装到resp的User对象中
		resp, err := etcd_discovery.UserClient.UserInfo(ctx, &user.UserInfoRequest{
			UserId: id,
			MyId:   myId,
		})
		if err != nil {
			return nil, err
		}
		resp2, err := etcd_discovery.MessageClient.GetLatestMessage(ctx, &message.MessageLatestRequest{
			MyId:   myId,
			UserId: id,
		})
		if err != nil {
			klog.Error("GetLatestMessage时发生了错误：" + err.Error())
		}
		friends.UserList = append(friends.UserList, &relation.FriendUser{
			Id:            resp.User.Id,
			Name:          resp.User.Name,
			FollowCount:   resp.User.FollowCount,
			FollowerCount: resp.User.FollowerCount,
			IsFollow:      resp.User.IsFollow,
			Avatar:        resp.User.Avatar,
			Message:       resp2.GetContent(),
			MsgType:       int64(resp2.GetMsgType()),
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

// MyList 实现排序
type MyList []*message.Message

// Len 实现sort.Interface接口的获取元素数量方法
func (m MyList) Len() int {
	return len(m)
}

// Less 实现sort.Interface接口的比较元素方法
func (m MyList) Less(i, j int) bool {
	return m[i].Id < m[j].Id
}

// Swap 实现sort.Interface接口的交换元素方法
func (m MyList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// GetFriendMessageList 获取所有朋友的聊天记录
func GetFriendMessageList(ctx context.Context, userId int64) (*relation.RelationFriendsMessageListResponse, error) {
	ids := model.GetFriendsByUserId(userId)
	messageList := &relation.RelationFriendsMessageListResponse{}
	var allList MyList
	for _, friendId := range ids {
		chat, err := etcd_discovery.MessageClient.GetMessageListByDB(ctx, &message.MessageChatRequest{
			UserId:   userId,
			ToUserId: friendId,
		})
		if err != nil {
			return nil, err
		}
		chat2, err := etcd_discovery.MessageClient.GetMessageListByDB(ctx, &message.MessageChatRequest{
			UserId:   friendId,
			ToUserId: userId,
		})
		//klog.Info("chat2:", chat2.MessageList)
		var myList MyList
		myList = append(chat.GetMessageList(), chat2.GetMessageList()...)
		sort.Sort(myList)
		allList = append(allList, myList...)

	}
	messageList = &relation.RelationFriendsMessageListResponse{
		MessageList: allList,
	}
	return messageList, nil
}
