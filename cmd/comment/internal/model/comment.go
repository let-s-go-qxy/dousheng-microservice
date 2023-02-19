package model

import (
	g "dousheng/pkg/global"
)

type Comment struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	UserId     int64  `json:"user_id"`
	Content    string `json:"comment_text" gorm:"column:comment_text"`
	CreateDate string `json:"create_time" gorm:"column:create_time"`
	VideoId    int64  `json:"video_id"`
	Cancel     int32  `gorm:"default:1" json:"cancel"`
}

// FindCommentByVideo 根据视频id查询全部评论，并且仅返回cancel等于1的
func FindCommentByVideo(videoId int64) []Comment {
	comments := make([]Comment, 0)
	g.MysqlDB.Where("video_id = ? AND cancel = 1 ", videoId).Order("create_time desc").Find(&comments)
	return comments
}

func GetCommentCount(videoId int64) (count int64) {
	g.MysqlDB.Model(&Comment{}).Where("video_id = ? AND cancel = ? ", videoId, 1).Count(&count)
	return
}

// FindCommentById 根据评论的id查询返回对应评论
func FindCommentById(id int64) (comment Comment) {
	g.MysqlDB.Where("id = ?", id).Find(&comment)
	return
}

// CreateComment comment表内插入对应评论
func CreateComment(comment *Comment) (returnComment *Comment, err error) {
	db := g.MysqlDB.Create(comment)
	model := db.Statement.Model
	returnComment = model.(*Comment)
	return
}

// DeleteComment comment表内删除对应评论,输入参数为comment整体
func DeleteComment(comment *Comment) (err error) {
	err = g.MysqlDB.Model(&Comment{}).Where("id =?", comment.Id).Update("cancel", "0").Error
	return
}
