package main

import (
	"context"
	videoService "dousheng/cmd/video/internal/service"
	video "dousheng/kitex_gen/video"
	g "dousheng/pkg/global"
	"github.com/jinzhu/copier"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	title := req.Title
	data := req.Data
	userId := req.UserId
	success := videoService.PublishVideo(int(userId), title, data)

	publishActionResponse := &video.PublishActionResponse{}

	if !success {
		publishActionResponse.StatusCode = g.StatusCodeFail
		publishActionResponse.StatusMsg = g.PublishVideoFailedMsg
	}

	publishActionResponse.StatusCode = g.StatusCodeOk
	publishActionResponse.StatusMsg = g.PublishVideoSuccessMsg

	return publishActionResponse, err
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	publishList, _, err := videoService.GetPublishList(int(req.UserId))

	var publishVideoListResp []*video.Video

	for _, publishVideo := range publishList {
		var publishVideoResp video.Video
		copier.Copy(&publishVideoResp, &publishVideo)

		author := publishVideo.Author
		respAuthor := *publishVideoResp.Author
		copier.Copy(&respAuthor, &author)

		publishVideoResp.Author = &respAuthor
		publishVideoListResp = append(publishVideoListResp, &publishVideoResp)
	}
	publishListResponse := &video.PublishListResponse{
		StatusCode: 0,
		StatusMsg:  "",
		VideoList:  publishVideoListResp,
	}
	return publishListResponse, err
}

// PublishVideoCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideoCount(ctx context.Context, req *video.PublishVideoCountRequest) (resp *video.PublishVideoCountResponse, err error) {
	_, messageCount, _ := videoService.GetPublishList(int(req.UserId))
	response := &video.PublishVideoCountResponse{}
	response.PublishVideoCount = int32(messageCount)
	return response, err
}

// GetFeedList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFeedList(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	latestTime := req.LatestTime
	userId := req.UserId
	nextTime, videoInfoList, state := videoService.GetVideoFeed(int64(latestTime), int32(userId))

	var videoListResp []*video.Video

	for _, videoInfo := range videoInfoList {
		v := video.Video{}
		copier.Copy(&v, &videoInfo)

		author := *v.Author
		copier.Copy(&author, videoInfo.Author)
		videoListResp = append(videoListResp, &v)
	}

	feedResponse := &video.FeedResponse{}
	feedResponse.StatusCode = g.StatusCodeOk
	feedResponse.StatusMsg = "获取视频Feed流成功！"
	feedResponse.VideoList = videoListResp
	feedResponse.NextTime = int32(nextTime)
	feedResponse.State = int32(state)
	return feedResponse, err
}
