package service

import (
	"context"
	"dousheng/cmd/video/internal/model"
	"dousheng/conf"
	"dousheng/kitex_gen/comment"
	"dousheng/kitex_gen/like"
	"dousheng/kitex_gen/user"
	"dousheng/kitex_gen/video"
	"dousheng/pkg/etcd_discovery"
	utils "dousheng/pkg/utils/file"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"
)

func GetVideoFeed(latestTime int64, userID int32) (nextTime int64, videoInfo []model.TheVideoInfo, state int) {
	// state 0:已经没有视频了  1:获取成功  -1:获取失败

	allVideoInfoData, isExist := model.VideoDao.GetVideoFeed(int32(latestTime))

	if !isExist {
		// 已经没有视频了
		return nextTime, videoInfo, 0
	}

	nextTime = int64(allVideoInfoData[len(allVideoInfoData)-1].Time)
	videoInfo = make([]model.TheVideoInfo, len(allVideoInfoData))

	wg := sync.WaitGroup{}
	wg.Add(len(allVideoInfoData))

	for index, videoInfoData := range allVideoInfoData {
		userInfoResponse, err := etcd_discovery.UserClient.UserInfo(context.Background(), &user.UserInfoRequest{
			MyId:   int64(userID),
			UserId: int64(videoInfoData.UserID),
		})
		if err != nil {
			klog.Error("调用UserInfo接口时发生了错误：" + err.Error())
		}

		commentCountResponse, err := etcd_discovery.CommentClient.CommentCount(context.Background(), &comment.CommentCountRequest{
			VideoId: int64(videoInfoData.VideoID),
		})
		if err != nil {
			klog.Error("调用CommentCount接口时发生了错误：" + err.Error())
		}

		favoriteCountResponse, err := etcd_discovery.LikeClient.FavoriteCount(context.Background(), &like.FavoriteCountRequest{
			VideoId: int64(videoInfoData.VideoID),
		})
		if err != nil {
			klog.Error("调用FavoriteCount接口时发生了错误：" + err.Error())
		}

		isFavoriteResponse, err := etcd_discovery.LikeClient.IsFavorite(context.Background(), &like.IsFavoriteRequest{
			UserId:  int64(userID),
			VideoId: int64(videoInfoData.VideoID),
		})
		if err != nil {
			klog.Error("调用IsFavorite接口时发生了错误：" + err.Error())
		}
		go func(index int, videoInfo []model.TheVideoInfo, videoInfoData model.VideoInfo, userID int32) {

			videoAuthorPointer := userInfoResponse.User
			videoAuthor := *videoAuthorPointer

			var followerCount, followCount, commentCount, vedioFavoriteCount int
			var isFollow, isFavorite bool
			followerCount = int(videoAuthor.GetFollowerCount())
			followCount = int(videoAuthor.GetFollowCount())
			commentCount = int(commentCountResponse.GetCommentCount())         //TODO
			vedioFavoriteCount = int(favoriteCountResponse.GetFavoriteCount()) //TODO
			isFollow = videoAuthor.GetIsFollow()
			isFavorite = isFavoriteResponse.GetIsFavorite() //TODO
			avatarURL := videoAuthor.Avatar
			workCount := videoAuthor.GetWorkCount()
			favoriteCount := videoAuthor.GetFavoriteCount()
			backgroundImage := videoAuthor.BackgroundImage
			signature := videoAuthor.Signature
			totalFavorite := videoAuthor.GetTotalFavorite()

			videoInfo[index] = model.TheVideoInfo{
				Id: videoInfoData.VideoID,
				Author: model.AuthorInfo{
					Id:              videoInfoData.UserID,
					Name:            videoInfoData.Username,
					FollowCount:     int(followCount),
					FollowerCount:   int(followerCount),
					IsFollow:        isFollow,
					Avatar:          avatarURL,
					WorkCount:       workCount,
					FavoriteCount:   favoriteCount,
					BackgroundImage: backgroundImage,
					Signature:       signature,
					TotalFavorite:   totalFavorite,
				},
				PlayUrl:       conf.OSSPreURL + videoInfoData.PlayURL + ".mp4",
				CoverUrl:      conf.OSSPreURL + videoInfoData.CoverURL + ".jpg",
				FavoriteCount: int(vedioFavoriteCount),
				CommentCount:  int(commentCount),
				IsFavorite:    isFavorite,
				Title:         videoInfoData.Title,
			}
			wg.Done()
		}(index, videoInfo, videoInfoData, userID)
		if err != nil {
			klog.Error("获取视频信息失败，出错了！")
			return nextTime, videoInfo, -1
		}
	}
	wg.Wait()
	return nextTime, videoInfo, 1
}

func PublishVideo(userID int, title string, fileBytes []byte) (success bool) {

	//雪花算法生成fileID
	node, _ := utils.NewWorker(1)
	randomId := node.GetId()
	fileID := fmt.Sprintf("%v", randomId)
	success = true

	if !utils.UploadFile(fileBytes, fileID, "video") {
		success = false
	}

	// 通过ffmpeg截取视频第一帧为视频封面
	videoURL := conf.OSSPreURL + fileID + ".mp4"
	//videoName := fileID + ".mp4"
	pictureName := fileID + ".jpg"

	//封面图和视频在本地的保存路径
	picturePath := conf.LocalFolderPath + pictureName
	//videoPath := ossRelated.LocalFolderPath + videoName

	//将上传的文件流的形式以mp4的形式保存到本地，并将视频的第一帧作为封面图导出到，picturePath下
	//ioutil.WriteFile(videoPath, fileBytes, 0666)
	cmd := exec.Command(conf.FfmpegPath, "-i", videoURL, "-y", "-f", "image2", "-ss", "1", picturePath)
	//buf := new(bytes.Buffer)
	//cmd.Stdout = buf
	cmd.Run()

	//以字节数组的形式读出本地的封面图，便于后续上传到云端
	openFile, err := os.Open(picturePath)
	defer openFile.Close()
	if err != nil {
		klog.Error("打开picture文件时发生了错误")
	}
	pictureBytes, err := ioutil.ReadAll(openFile)
	if err != nil {
		klog.Error("读取picture文件时发生了错误")
	}

	// 将视频封面上传至OSS中
	if !utils.UploadFile(pictureBytes, fileID, "picture") {
		success = false
	}
	if success {
		if model.VideoDao.PublishVideo(userID, title, fileID) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// 获取登录用户的视频发布列表
func GetPublishList(userId, myId int) (respVideoList []model.RespVideo, publishCount int, err error) {

	//获取视频数组
	var videoList []model.Video
	videoList = model.GetPublishList(userId)

	//利用封装函数
	respVideoList = PlusAuthor(userId, myId, videoList)
	publishCount = len(respVideoList)
	return
}

func GetPublishVideoCount(userId int) (publishCount int) {
	var videoList []model.Video
	videoList = model.GetPublishList(userId)
	publishCount = len(videoList)
	return publishCount
}

// 将author封装到video
func PlusAuthor(userId, myId int, videoList []model.Video) (respVideoList []model.RespVideo) {

	for _, video := range videoList {
		respVideo := model.RespVideo{}
		copier.Copy(&respVideo, &video)
		author := model.Author{}

		author.Id = int(video.Author)

		userInfoResponse, err := etcd_discovery.UserClient.UserInfo(context.Background(), &user.UserInfoRequest{
			UserId: int64(userId),
			MyId:   int64(myId),
		})
		if err != nil {
			klog.Error("调用UserInfo接口时发生了错误：" + err.Error())
		}

		commentCountResponse, err := etcd_discovery.CommentClient.CommentCount(context.Background(), &comment.CommentCountRequest{
			VideoId: int64(video.Id),
		})
		if err != nil {
			klog.Error("调用CommentCount接口时发生了错误：" + err.Error())
		}

		favoriteCountResponse, err := etcd_discovery.LikeClient.FavoriteCount(context.Background(), &like.FavoriteCountRequest{
			VideoId: int64(video.Id),
		})
		if err != nil {
			klog.Error("调用FavoriteCount接口时发生了错误：" + err.Error())
		}

		isFavoriteResponse, err := etcd_discovery.LikeClient.IsFavorite(context.Background(), &like.IsFavoriteRequest{
			UserId:  int64(userId),
			VideoId: int64(video.Id),
		})
		if err != nil {
			klog.Error("调用IsFavorite接口时发生了错误：" + err.Error())
		}
		videoAuthorPointer := userInfoResponse.User
		videoAuthor := *videoAuthorPointer

		author.FollowCount = int(videoAuthor.FollowCount)
		author.FollowerCount = int(videoAuthor.FollowerCount)
		copier.Copy(&author, &videoAuthor)

		copier.Copy(&respVideo.Author, &author)
		respVideo.Id = int(video.Id)
		respVideo.PlayUrl = conf.OSSPreURL + respVideo.PlayUrl + ".mp4"
		respVideo.CoverUrl = conf.OSSPreURL + respVideo.CoverUrl + ".jpg"

		respVideo.CommentCount = int(commentCountResponse.GetCommentCount())
		respVideo.FavoriteCount = int(favoriteCountResponse.GetFavoriteCount())
		respVideo.IsFavorite = isFavoriteResponse.GetIsFavorite()

		respVideoList = append(respVideoList, respVideo)
	}
	return
}

// GetPublishIds 获取该用户发表的视频id数组
func GetPublishIds(c context.Context, userID int64) (resp *video.PublishIdsResponse, err error) {
	return &video.PublishIdsResponse{
		Ids: model.FindVideoIds(userID),
	}, nil
}

// GetVideoInfo  获取视频详情
func GetVideoInfo(c context.Context, videoId int64) (resp *video.VideoInfoResponse, err error) {
	userInfo := &video.User{}
	videoInfo, err := model.FindVideoById(videoId)
	if err != nil {
		return
	}
	userInfo2, err := etcd_discovery.UserClient.UserInfo(c, &user.UserInfoRequest{
		UserId: int64(videoInfo.Author),
		MyId:   0,
	})
	isLikeResp, err := etcd_discovery.LikeClient.IsFavorite(c, &like.IsFavoriteRequest{
		UserId:  int64(videoInfo.Author),
		VideoId: videoId,
	})
	favoriteCountResp, err := etcd_discovery.LikeClient.FavoriteCount(c, &like.FavoriteCountRequest{
		VideoId: videoId,
	})
	if err != nil {
		return
	}
	copier.Copy(userInfo, userInfo2)
	return &video.VideoInfoResponse{
		VideoInfo: &video.Video{
			Id:            int64(videoInfo.Id),
			Author:        userInfo,
			PlayUrl:       videoInfo.PlayUrl,
			CoverUrl:      videoInfo.CoverUrl,
			FavoriteCount: favoriteCountResp.GetFavoriteCount(),
			CommentCount:  0, // 不用就不给了
			IsFavorite:    isLikeResp.GetIsFavorite(),
			Title:         videoInfo.Title,
		},
	}, nil
}
