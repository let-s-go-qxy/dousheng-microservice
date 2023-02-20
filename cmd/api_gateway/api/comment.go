package api

import (
	"context"
	"dousheng/kitex_gen/comment"
	"dousheng/pkg/etcd_discovery"
	g "dousheng/pkg/global"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
)

// GetCommentList 获取视频id的评论，以评论时间排序
func GetCommentList(c context.Context, ctx *app.RequestContext) {
	// 获取请求参数
	videoId := ctx.Query("video_id")
	userIDInterface, success := ctx.Get("user_id")
	var userId int64
	if success {
		userId = userIDInterface.(int64)
	} // 若不存在，userID默认为0
	vid, err := strconv.Atoi(videoId)
	if err != nil {
		klog.Error("获取视频id失败")
	}
	req := &comment.CommentListRequest{
		UserId:  int64(userId),
		VideoId: int64(vid),
	}

	resp, _ := etcd_discovery.CommentClient.GetCommentList(c, req)
	commentList := resp.GetCommentList()
	for _, comment := range commentList {
		comment.CreateDate = comment.CreateDate[5:10]
	}
	ctx.JSON(consts.StatusOK, resp)
}

// PostCommentAction 对视频下的评论进行发表或者删除
func PostCommentAction(c context.Context, ctx *app.RequestContext) {
	// 获取请求参数
	value, _ := ctx.Get("user_id")
	userId := value.(int64)
	videoId, _ := strconv.Atoi(ctx.Query("video_id"))       //》根据视频查找对应评论
	actionType, _ := strconv.Atoi(ctx.Query("action_type")) //》视频操作？1》添加insert：2》删除delete
	commentText := ctx.Query("comment_text")
	commentId, _ := strconv.Atoi(ctx.Query("comment_id"))

	req := &comment.CommentActionRequest{
		UserId:      int64(userId),
		VideoId:     int64(videoId),
		ActionType:  int32(actionType),
		CommentText: commentText,
		CommentId:   int64(commentId),
	}

	resp, err := etcd_discovery.CommentClient.PostCommentAction(c, req)

	if err != nil {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  err.Error(),
		})
	}

	ctx.JSON(consts.StatusOK, resp)
}
