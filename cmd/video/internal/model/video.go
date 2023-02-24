package model

import (
	g "dousheng/pkg/global"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"
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
	Id              int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	FollowCount     int    `json:"follow_count,omitempty"`
	FollowerCount   int    `json:"follower_count,omitempty"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	WorkCount       int32  `json:"work_count"`
	FavoriteCount   int32  `json:"favorite_count"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorite   int32  `json:"total_favorite"`
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
	Id              int32  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int    `json:"follow_count"`
	FollowerCount   int    `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	WorkCount       int32  `json:"work_count"`
	FavoriteCount   int32  `json:"favorite_count"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorite   int32  `json:"total_favorite"`
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
	g.WriteMysqlDB.Table("videos").Debug().Create(&video)
	return true

}

// 通过用户id查询video数组
func GetPublishList(userId int) (videoList []Video) {

	g.ReadMysqlDB.Table("videos").
		Where("author_id= ? ", userId).
		Scan(&videoList)

	return
}

func (*VideoDaoStruct) GetVideoFeed(latestTime int32) ([]VideoInfo, bool) {

	var result []VideoInfo

	g.ReadMysqlDB.Debug().Raw("SELECT `users`.`id` AS `UserID`,`users`.`name` AS `Username`, `videos`.`id` AS `VideoID`,"+
		"`videos`.`play_url`, `videos`.`cover_url`,`videos`.`publish_time` AS `Time`,`videos`.`title` "+
		"FROM `videos` INNER JOIN `users` ON `users`.`id` = `videos`.`author_id` "+
		"WHERE `videos`.`publish_time` < ? ORDER BY `videos`.`publish_time` DESC LIMIT 10", latestTime).Scan(&result)

	if result == nil {
		return nil, false
	}
	return result, true
}

func FindVideoIds(userId int64) (ids []int64) {
	videos := []Video{}
	g.ReadMysqlDB.Select("id").Find(&videos, "author_id = ?", userId)
	for _, video := range videos {
		ids = append(ids, int64(video.Id))
	}
	return
}

func FindVideoById(videoId int64) (Video, error) {
	video := &Video{
		Id: int32(videoId),
	}
	err := g.ReadMysqlDB.First(&video).Error
	return *video, err
}

func GetVideoCache(videoId int32) *TheVideoInfo {
	result, err := g.DbVerify.Get(g.RedisContext, strconv.Itoa(int(videoId))).Result()
	if err != nil {
		hlog.Error("GetVideoCache err: ", err.Error())
		return nil
	}
	theVideoInfo := &TheVideoInfo{}
	err = json.Unmarshal([]byte(result), theVideoInfo)
	if err != nil {
		return nil
	}
	return theVideoInfo
}

// CacheVideo 判断是否是热门视频并存储
func CacheVideo(video TheVideoInfo) {
	// TODO 这里判断热门视频的策略有待优化，目前判断之前访问过100次就算是热门视频
	g.DbVerify.Incr(g.RedisContext, strconv.Itoa(int(video.Id))+"count")
	var val string
	val, _ = g.DbVerify.Get(g.RedisContext, strconv.Itoa(int(video.Id))+"count").Result()
	count, _ := strconv.Atoi(val)
	if count > 100 {
		bits, _ := json.Marshal(video)
		g.DbVerify.Set(g.RedisContext, strconv.Itoa(int(video.Id)), string(bits), time.Minute*100)
		g.DbVerify.Set(g.RedisContext, strconv.Itoa(int(video.Id))+"count", "0", time.Minute*10)
	}
}
