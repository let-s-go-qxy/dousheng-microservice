package api

import (
	"context"
	"dousheng/kitex_gen/like"
	"dousheng/pkg/etcd_discovery"
	"dousheng/pkg/utils/msg"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
)

// GetFavoriteList 获取喜爱视频列表
func GetFavoriteList(c context.Context, ctx *app.RequestContext) {
	userId := ctx.Query("user_id")
	strUserId, err := strconv.Atoi(userId)
	if err != nil {
		klog.Error("用户ID错误")
	}
	req := &like.FavoriteListRequest{
		UserId: int64(strUserId),
		MyId:   int64(strUserId),
	}

	//videoList, _ := like.GetFavoriteList(strUserId)
	resp, _ := etcd_discovery.LikeClient.GetFavoriteList(c, req)
	videoList := resp.VideoList
	respVideoList := make([]Video, len(videoList))
	copier.Copy(&respVideoList, &videoList)
	favoriteListResponse := FavoriteListResponse{StatusCode: 0, StatusMsg: "返回成功", VideoList: respVideoList}
	ctx.JSON(consts.StatusOK, favoriteListResponse)
}

// FavoriteAction 点赞和取消点赞操作
func FavoriteAction(c context.Context, ctx *app.RequestContext) {
	value, _ := ctx.Get("user_id")
	userId := value.(int)
	videoId, _ := strconv.Atoi(ctx.Query("video_id"))
	actionType, _ := strconv.Atoi(ctx.Query("action_type"))

	req := &like.FavoriteActionRequest{
		UserId:     int64(userId),
		VideoId:    int64(videoId),
		ActionType: int32(actionType),
	}

	_, err := etcd_discovery.LikeClient.FavoriteAction(c, req)
	if err == nil {
		ctx.JSON(consts.StatusOK,
			Response{0,
				msg.LikeFavoriteActionSuccess})
	} else {
		ctx.JSON(consts.StatusOK, Response{1,
			msg.LikeFavoriteActionFail})
	}
}
