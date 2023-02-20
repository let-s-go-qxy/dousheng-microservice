package model

import (
	g "dousheng/pkg/global"
	"gorm.io/gorm"
)

type Follow struct {
	Id       int64 `gorm:"primaryKey" json:"id"`
	UserId   int64 `json:"user_id"`
	FollowId int64 `json:"follow_id"`
	Cancel   int   `json:"cancel"`
}

// Author 用户返回模型
type Author struct {
	Id            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int    `json:"follow_count,omitempty"`
	FollowerCount int    `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow"`
	Avatar        string `json:"avatar"`
}

// GetFollowsByUserId  查所有的被关注者的id
func GetFollowsByUserId(userId int64) (arr []int64) {
	follows := new([]Follow)
	g.MysqlDB.Find(follows, "user_id = ? AND cancel = ?", userId, 1)
	for _, follow := range *follows {
		arr = append(arr, follow.FollowId)
	}
	return
}

// GetFollowersByUserId 获取所有粉丝的id
func GetFollowersByUserId(userId int64) (arr []int64) {
	follows := new([]Follow)
	g.MysqlDB.Find(follows, "follow_id = ? AND cancel = ?", userId, 1)
	for _, follow := range *follows {
		arr = append(arr, follow.UserId)
	}
	return
}

func GetFriendsByUserId(userId int64) (arr []int64) {
	follows := new([]Follow)
	g.MysqlDB.Find(follows, "user_id = ? AND cancel = ?", userId, 1)
	for _, follow := range *follows {
		var count int64
		g.MysqlDB.Model(&Follow{}).
			Where("user_id = ? AND follow_id = ? AND cancel = ?", follow.FollowId, userId, 1).
			Count(&count)
		if count > 0 {
			arr = append(arr, follow.FollowId)
		}
	}
	return
}

// GetFollowCount 获取当前用户的关注人数
func GetFollowCount(userId int64) (count int64) {
	g.MysqlDB.Model(&Follow{}).Where("user_id = ? AND cancel = ?", userId, 1).Count(&count)
	return
}

// GetFollowerCount 获取当前用户的粉丝人数
func GetFollowerCount(userId int64) (count int64) {
	g.MysqlDB.Model(&Follow{}).Where("follow_id = ? AND cancel = ?", userId, 1).Count(&count)
	return
}

func IsFollow(userId, followId int64) bool {
	err := g.MysqlDB.Debug().First(&Follow{}, "user_id = ? AND follow_id = ? AND cancel = ?", userId, followId, 1).Error
	return err == nil
}

// CreateOrUpdateFollow 新增或更新记录
func CreateOrUpdateFollow(myId, userId int64, followType int32) error {
	follow := new(Follow)
	// 如果有记录更新记录，没有新增记录
	if err := g.MysqlDB.First(follow, "user_id = ? AND follow_id = ?", myId, userId).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		follow.UserId = myId
		follow.FollowId = userId
		follow.Cancel = int(followType)
		return g.MysqlDB.Create(follow).Error
	}
	return g.MysqlDB.Model(follow).
		Where("user_id = ? AND follow_id = ?", myId, userId).
		Update("cancel", followType).Error
}
