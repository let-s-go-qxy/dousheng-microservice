package api

import (
	"context"
	"dousheng/kitex_gen/video"
	"dousheng/pkg/etcd_discovery"
	g "dousheng/pkg/global"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// AuthorInfo 作者信息
type AuthorInfo struct {
	ID            int32  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	Avatar        string `json:"avatar"`
}

// TheVideoInfo 视频信息
type TheVideoInfo struct {
	ID            int32      `json:"id"`
	Author        AuthorInfo `json:"author"`
	PlayUrl       string     `json:"play_url"`
	CoverUrl      string     `json:"cover_url"`
	FavoriteCount int        `json:"favorite_count"`
	CommentCount  int        `json:"comment_count"`
	IsFavorite    bool       `json:"is_favorite"`
	Title         string     `json:"title"`
}

type GetVideoResponse struct {
	g.Response
	NextTime  int64          `json:"next_time"`
	VideoList []TheVideoInfo `json:"video_list"`
}

func GetFeedList(c context.Context, ctx *app.RequestContext) {
	latestTime, _ := strconv.ParseInt(ctx.Query("latest_time"), 10, 32)
	userIDInterface, success := ctx.Get("user_id")
	var userID int32
	if success {
		userID = int32(userIDInterface.(int64))
	} // 若不存在，userID默认为0

	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}

	feedRequest := &video.FeedRequest{}
	feedRequest.UserId = int64(userID)
	feedRequest.LatestTime = int32(latestTime)

	// 需要获取NextTime、VideoList
	getFeedListResp, err := etcd_discovery.VideoClient.GetFeedList(c, feedRequest)
	if err != nil {
		klog.Error("GetFeedList时发生了错误：" + err.Error())
	}
	nextTime := getFeedListResp.NextTime
	videoInfoList := getFeedListResp.VideoList
	state := getFeedListResp.State
	//nextTime, videoInfo, state := video.GetVideoFeed(latestTime, userID)

	var videoInfoListResp []TheVideoInfo
	var videoInfoResp TheVideoInfo
	for _, videoInfo := range videoInfoList {

		v := *videoInfo
		videoInfoResp.ID = int32(v.Id)
		copier.Copy(&videoInfoResp, &v)

		vAuthor := *v.Author
		videoInfoResp.Author.ID = int32(vAuthor.Id)

		videoInfoListResp = append(videoInfoListResp, videoInfoResp)
	}

	if state == g.FeedStatusNull {
		ctx.JSON(http.StatusOK, &GetVideoResponse{
			Response: g.Response{
				StatusCode: g.StatusCodeFail,
				StatusMsg:  g.HasNoVideoMsg,
			}, NextTime: latestTime,
		})
	} else if state == g.FeedStatusError {
		ctx.JSON(http.StatusOK, &GetVideoResponse{
			Response: g.Response{
				StatusCode: g.StatusCodeFail,
				StatusMsg:  g.GetVideoInfoErrorMsg,
			}, NextTime: latestTime,
		})
	} else if state == g.StatusCodeOk {

		ctx.JSON(http.StatusOK, &GetVideoResponse{
			Response: g.Response{
				StatusCode: g.StatusCodeOk,
				StatusMsg:  g.GetVideoInfoSuccessMsg,
			}, NextTime: int64(nextTime),
			VideoList:   videoInfoListResp,
		})
	}
}

func PublishVideo(c context.Context, ctx *app.RequestContext) {
	title := ctx.PostForm("title")
	data, err := ctx.FormFile("data")
	userID, success := ctx.Get("user_id")
	if !success {
		ctx.JSON(http.StatusOK,
			g.Response{
				StatusCode: -1,
				StatusMsg:  g.TokenParameterAcquisitionError,
			})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK,
			g.Response{
				StatusCode: -1,
				StatusMsg:  g.PublishVideoFailedMsg,
			})
		return
	}

	fileHandle, err1 := data.Open() //打开上传文件
	if err1 != nil {
		klog.Error("打开文件失败" + err1.Error())
	}

	// 闭包处理错误
	defer func(fileHandle multipart.File) {
		err := fileHandle.Close()
		if err != nil {
			klog.Error("关闭文件错误" + err.Error())
		}
	}(fileHandle)

	fileByte, err2 := ioutil.ReadAll(fileHandle)
	if err2 != nil {
		klog.Error("读取文件错误" + err2.Error())
	}

	publishVideoReq := &video.PublishActionRequest{}
	publishVideoReq.Title = title
	publishVideoReq.UserId = userID.(int64)
	publishVideoReq.Data = fileByte

	publishVideoResponse, err := etcd_discovery.VideoClient.PublishVideo(c, publishVideoReq)

	if publishVideoResponse.StatusCode == g.StatusOk {
		ctx.JSON(http.StatusOK,
			g.Response{
				StatusCode: g.StatusOk,
				StatusMsg:  g.PublishVideoSuccessMsg,
			})
	} else {
		ctx.JSON(http.StatusOK,
			g.Response{
				StatusCode: g.StatusCodeFail,
				StatusMsg:  g.PublishVideoFailedMsg,
			})
	}
}

// PublishList 发布列表
func PublishList(c context.Context, ctx *app.RequestContext) {

	userId, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		klog.Error("用户ID错误")
	}

	publishListRequest := &video.PublishListRequest{}
	publishListRequest.UserId = int64(userId)

	publishListResp, err := etcd_discovery.VideoClient.PublishList(c, publishListRequest)
	videoList := publishListResp.VideoList

	var publishVideoListResp []Video
	var publishVideoResp Video
	for _, publishVideoPointer := range videoList {
		v := *publishVideoPointer
		copier.Copy(&publishVideoResp, &v)
		publishVideoListResp = append(publishVideoListResp, publishVideoResp)
	}

	resp := VideoListResponse{Response: Response{
		StatusCode: g.StatusCodeOk, StatusMsg: "成功!!"},
		VideoList: publishVideoListResp}
	ctx.JSON(consts.StatusOK, resp)
}
