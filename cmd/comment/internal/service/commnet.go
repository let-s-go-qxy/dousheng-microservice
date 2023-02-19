package service

import (
	"dousheng/cmd/comment/internal/model"
	"dousheng/kitex_gen/comment"
	"errors"
	"time"

	"github.com/jinzhu/copier"
)

// GetCommentList 查选该视频下的所有评论
func GetCommentList(videoId int64, myId int64) (respCommentList []comment.Comment) {

	// 调用model层comment的sql查询语句，根据视频id查询对应id的视频评论
	commentList := model.FindCommentByVideo(videoId)
	for _, comments := range commentList {
		respComment := comment.Comment{}
		copier.Copy(respComment, comments)
		respCommentList = append(respCommentList, respComment)
	}

	return
}

// CommentAction 对评论进行创建或者删除
func CommentAction(videoId int64, actionType int32, contentText string, commentId int64, userId int64) (id int64, content string, createTime string, err error) {
	// 填装Comment数据
	com := &model.Comment{
		Id:         commentId,
		UserId:     userId,
		Content:    contentText,
		CreateDate: time.Now().Format("2006-01-02 15:04:05"),
		VideoId:    videoId,
		Cancel:     actionType,
	}
	if actionType == 1 {
		com, err = model.CreateComment(com)
		// 评论创建失败
		if err != nil {
			err = errors.New("发表评论失败: " + err.Error())

		}
	} else {
		err = model.DeleteComment(com)
		if err != nil {
			err = errors.New("删除评论失败: " + err.Error())
		}
	}
	return com.Id, com.Content, com.CreateDate, err
}

func GetCommentCount(videoId int64) (commentCount int32) {
	commentCount = int32(model.GetCommentCount(videoId))
	return
}
