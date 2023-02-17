package api

import (
	"context"
	"dousheng/pkg/etcd_discovery"
	g "dousheng/pkg/global"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	"strconv"
)

// GetCommentList 获取视频id的评论，以评论时间排序
func GetCommentList(c context.Context, ctx *app.RequestContext) {
	// 获取请求参数
	videoId := ctx.Query("video_id")
	userIDInterface, success := ctx.Get("user_id")
	var userId int32
	if success {
		userId = int32(userIDInterface.(int))
	} // 若不存在，userID默认为0
	vid, err := strconv.Atoi(videoId)
	if err != nil {

	}
	// 1、通过video表查询对应主键服务；2、根据video主键作为查询条件，查询相应评论
	comments, videoCommentCount := etcd_discovery.CommentClient.GetCommentList(vid, int(userId))
	print(videoCommentCount)
	respCommentList := make([]Comment, 0)
	for _, comment := range comments {
		respComment := Comment{}
		copier.Copy(&respComment, &comment)
		respComment.CreateTime = respComment.CreateTime[5:10]
		respCommentList = append(respCommentList, respComment)
	}
	resp := CommentListResponse{StatusCode: 0, StatusMsg: "返回成功", CommentList: respCommentList}
	ctx.JSON(consts.StatusOK, resp)
}

// PostCommentAction 对视频下的评论进行发表或者删除
func PostCommentAction(c context.Context, ctx *app.RequestContext) {
	// 获取请求参数
	videoId, _ := strconv.Atoi(ctx.Query("video_id"))       //》根据视频查找对应评论
	actionType, _ := strconv.Atoi(ctx.Query("action_type")) //》视频操作？1》添加insert：2》删除delete
	commentText := ctx.Query("comment_text")
	commentId, _ := strconv.Atoi(ctx.Query("comment_id"))
	// 获取userId
	value, _ := ctx.Get("user_id")
	userId := value.(int)
	// 进行评论修改
	comment, userDao, err := sc.CommentAction(videoId, actionType, commentText, commentId, userId)
	if err != nil {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  err.Error(),
		})
	}
	respComment := Comment{}
	respComment.Id = comment.Id
	copier.Copy(&respComment.User, &userDao)
	respComment.Content = comment.Content
	respComment.CreateTime = comment.CreateTime[5:10]
	resp := CommentResponse{StatusCode: 0, StatusMsg: "返回成功", Comment: respComment}
	ctx.JSON(consts.StatusOK, resp)
}
