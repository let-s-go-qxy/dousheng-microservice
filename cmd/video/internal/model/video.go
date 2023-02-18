package model

import (
	g "dousheng/pkg/global"
	"sync"
	"time"
)

// Video mapped from table <video>
type Video struct {
	Id          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Author      int32  `gorm:"column:author_id;not null" json:"author"` //用户Id author_id
	PlayUrl     string `gorm:"column:play_url;not null" json:"play_url"`
	CoverUrl    string `gorm:"column:cover_url;not null" json:"cover_url"`
	PublishTime int64  `gorm:"column:publish_time;not null" json:"time"`
	Title       string `gorm:"column:title;not null" json:"title"`
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

// RespVideo 喜爱的视频返回模型
type RespVideo struct {
	Id            int    `json:"id,omitempty"`
	Author        Author `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int    `json:"favorite_count,omitempty"`
	CommentCount  int    `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

// TheVideoInfo 视频信息
type TheVideoInfo struct {
	Id            int32      `json:"id"`
	Author        AuthorInfo `json:"author"`
	PlayUrl       string     `json:"play_url"`
	CoverUrl      string     `json:"cover_url"`
	FavoriteCount int        `json:"favorite_count"`
	CommentCount  int        `json:"comment_count"`
	IsFavorite    bool       `json:"is_favorite"`
	Title         string     `json:"title"`
}

// AuthorInfo 作者信息
type AuthorInfo struct {
	Id            int32  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	Avatar        string `json:"avatar"`
}

type GetVideoResponse struct {
	g.Response
	NextTime  int64          `json:"next_time"`
	VideoList []TheVideoInfo `json:"video_list"`
}
type VideoInfo struct {
	VideoID       int32
	UserID        int32
	Username      string
	PlayURL       string
	CoverURL      string
	FavoriteCount int
	IsFavorite    bool
	Time          int32
	Title         string
}

type VideoDaoStruct struct {
}

var (
	VideoDao  *VideoDaoStruct
	videoOnce sync.Once
)

func init() {
	videoOnce.Do(func() {
		VideoDao = &VideoDaoStruct{}
	})
}

func (*VideoDaoStruct) PublishVideo(userID int, title string, videoNumID string) bool {
	video := Video{
		Author:      int32(userID),
		PlayUrl:     videoNumID,
		CoverUrl:    videoNumID,
		Title:       title,
		PublishTime: time.Now().Unix(),
	}
	g.MysqlDB.Table("videos").Debug().Create(&video)
	return true

}

// 通过用户id查询video数组
func GetPublishList(userId int) (videoList []Video) {

	g.MysqlDB.Table("videos").
		Where("author_id= ? ", userId).
		Scan(&videoList)

	return
}

func (*VideoDaoStruct) GetVideoFeed(latestTime int32) ([]VideoInfo, bool) {

	var result []VideoInfo

	g.MysqlDB.Debug().Raw("SELECT `users`.`id` AS `UserID`,`users`.`name` AS `Username`, `videos`.`id` AS `VideoID`,"+
		"`videos`.`play_url`, `videos`.`cover_url`,`videos`.`publish_time` AS `Time`,`videos`.`title` "+
		"FROM `videos` INNER JOIN `users` ON `users`.`id` = `videos`.`author_id` "+
		"WHERE `videos`.`publish_time` < ? ORDER BY `videos`.`publish_time` DESC LIMIT 10", latestTime).Scan(&result)

	if result == nil {
		return nil, false
	}
	return result, true
}
