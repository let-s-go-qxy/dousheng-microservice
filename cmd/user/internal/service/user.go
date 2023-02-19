package service

import (
	"context"
	"dousheng/cmd/user/internal/model"
	"dousheng/conf"
	"dousheng/kitex_gen/relation"
	"dousheng/kitex_gen/user"
	"dousheng/kitex_gen/video"
	"dousheng/pkg/etcd_discovery"
	g "dousheng/pkg/global"
	utils2 "dousheng/pkg/utils"
	"dousheng/pkg/utils/file"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// JudgeNameAndPassword 判断用户名和密码是否符合要求
func JudgeNameAndPassword(name, password string) bool {
	if len(name) > 32 || len(password) > 32 || name == "" || len(password) < 6 {
		return false
	}
	return true
}

func UserLogin(name, password string) (userId int64, token string, err error) {
	if JudgeNameAndPassword(name, password) == false {
		err = errors.New("账号或密码不符合要求")
		return
	}
	user := &model.User{
		Name: name,
	}
	err = model.GetUser(user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.New("账号或密码错误")
			return
		}
		err = errors.New("查询失败: " + err.Error())
		return
	}
	if utils2.GetMd5Str(password+user.Salt) != user.Password {
		err = errors.New("账号或密码错误")
		return
	}
	userId = user.Id
	token = GenerateToken(*user)
	return
}

func UserRegister(name, password string) (userId int64, token string, err error) {
	if JudgeNameAndPassword(name, password) == false {
		err = errors.New("账号或密码不符合要求")
		return
	}
	salt := strconv.Itoa(int(time.Now().UnixNano()))
	password = utils2.GetMd5Str(password + salt)
	// 填装数据
	user := &model.User{
		Name:     name,
		Password: password,
		Salt:     salt,
	}
	_, err = model.CreateUser(user)
	// 创建失败
	if err != nil {
		if err == g.ErrDbCreateUniqueKeyRepeatedly {
			err = errors.New("User already exist")
			return
		}
		err = errors.New("创建用户失败: " + err.Error())
		return
	}
	userId = int64(user.Id)

	//注册成功时自动生成一张用户头像到阿里OSS云端
	success := file.UploadAvatar(int(userId))
	if !success {
		klog.Error("上传用户头像失败！")
	}

	//注册成功时自动生成一张背景图到阿里OSS云端
	success = file.UploadBackground(int(userId))
	if !success {
		klog.Error("上传用户头像失败！")
	}

	token = GenerateToken(*user)
	return
}

func UserInfo(myId int64, userId int64) (userInfo user.User, err error) {
	userDao := &model.User{
		Id: userId,
	}
	err = model.GetUser(userDao)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.New("user not exit")
		}
		return
	}
	userInfo = user.User{
		Id:              userId,
		Name:            userDao.Name,
		BackgroundImage: GetBackgroundImage(userId),
		Avatar:          GetAvatar(userId),
		Signature:       "我不想再当一个xx了，我只想拥有快乐（谢谢你，狗子）",
	}
	// 获取用户关注数
	resp, err := etcd_discovery.RelationClient.GetFollowCount(context.Background(), &relation.RelationFollowCountRequest{
		UserId: userId,
	})
	if err != nil {
		return
	}
	userInfo.FollowCount = resp.GetCount()
	// 获取用户粉丝数
	resp2, err := etcd_discovery.RelationClient.GetFollowerCount(context.Background(), &relation.RelationFollowerCountRequest{
		UserId: userId,
	})
	if err != nil {
		return
	}
	userInfo.FollowerCount = resp2.GetCount()
	// 获取用户是否被关注
	resp3, err := etcd_discovery.RelationClient.GetIsFollow(context.Background(), &relation.RelationIsFollowRequest{
		MyId:   myId,
		UserId: userId,
	})
	userInfo.IsFollow = resp3.GetIsFollow()
	// 获取用户作品数
	resp4, _ := etcd_discovery.VideoClient.PublishVideoCount(context.Background(), &video.PublishVideoCountRequest{
		UserId: userId,
	})
	userInfo.WorkCount = resp4.GetPublishVideoCount()
	// TODO 喜欢数
	return
}

// GetAvatar 获取用户头像
func GetAvatar(userID int64) string {
	strUserID := strconv.Itoa(int(userID))
	avatarURL := conf.OSSAvatarPreURL + strUserID + "_avatar.jpg"
	return avatarURL
}

// GetBackgroundImage 获取用户背景图
func GetBackgroundImage(userID int64) string {
	strUserID := strconv.Itoa(int(userID))
	backgroundURL := conf.OSSBackgroundPreURL + strUserID + "_background.jpg"
	return backgroundURL
}
