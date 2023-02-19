package file

import (
	"bytes"
	g "dousheng/pkg/global"
	"github.com/cloudwego/kitex/pkg/klog"
	"io/ioutil"
	"os"
	"strconv"
)

func UploadFile(file []byte, filename string, fileType string) bool {
	var fileSuffix string
	if fileType == "video" {
		fileSuffix = ".mp4"
	} else if fileType == "picture" {
		fileSuffix = ".jpg"
	} else {
		klog.Error("无法上传" + fileType + "类型文件")
		return false
	}
	//初始化bucket

	err := g.OssBucket.PutObject("video/"+filename+fileSuffix, bytes.NewReader(file))
	if err != nil {
		klog.Error("上传文件失败" + err.Error())
		return false
	} else {
		return true
	}
}

func UploadAvatar(userID int) bool {

	var fileSuffix string
	fileSuffix = ".jpg"
	strUserID := strconv.Itoa(userID)

	ModUserID := userID % 10
	strModUserID := strconv.Itoa(ModUserID)
	//项目中头像相对路径
	avatarPath := "imgs/" +
		strModUserID +
		"_avatar.jpg"

	//以字节数组的形式读出本地的头像，便于后续上传到云端
	openFile, err := os.Open(avatarPath)
	defer openFile.Close()
	avatarBytes, err := ioutil.ReadAll(openFile)
	if err != nil {
		klog.Error("上传头像失败" + err.Error())
		return false
	}
	filename := strUserID + "_avatar"
	err = g.OssBucket.PutObject("avatar/"+filename+fileSuffix, bytes.NewReader(avatarBytes))
	if err != nil {
		klog.Error("上传头像失败" + err.Error())
		return false
	} else {
		return true
	}
}

func UploadBackground(userID int) bool {

	var fileSuffix string
	fileSuffix = ".jpg"
	strUserID := strconv.Itoa(userID)

	ModUserID := userID % 10
	strModUserID := strconv.Itoa(ModUserID)
	//项目中头像相对路径
	avatarPath := "imgs/" +
		strModUserID +
		"_background.jpg"

	//以字节数组的形式读出本地的头像，便于后续上传到云端
	openFile, err := os.Open(avatarPath)
	defer openFile.Close()
	avatarBytes, err := ioutil.ReadAll(openFile)
	if err != nil {
		klog.Error("上传头像失败" + err.Error())
		return false
	}
	filename := strUserID + "_background"
	err = g.OssBucket.PutObject("background/"+filename+fileSuffix, bytes.NewReader(avatarBytes))
	if err != nil {
		klog.Error("上传头像失败" + err.Error())
		return false
	} else {
		return true
	}
}
