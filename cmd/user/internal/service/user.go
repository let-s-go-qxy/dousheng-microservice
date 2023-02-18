package service

import (
	"context"
	"dousheng/cmd/user/internal/model"
	"dousheng/conf"
	"dousheng/kitex_gen/relation"
	"dousheng/kitex_gen/user"
	"dousheng/pkg/etcd_discovery"
	g "dousheng/pkg/global"
	utils2 "dousheng/pkg/utils"
	"errors"
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
	resp, err := etcd_discovery.RelationClient.GetFollowCount(context.Background(), &relation.RelationFollowCountRequest{
		UserId: userId,
	})
	if err != nil {
		return
	}
	userInfo.FollowCount = resp.GetCount()
	resp2, err := etcd_discovery.RelationClient.GetFollowerCount(context.Background(), &relation.RelationFollowerCountRequest{
		UserId: userId,
	})
	if err != nil {
		return
	}
	userInfo.FollowerCount = resp2.GetCount()
	resp3, err := etcd_discovery.RelationClient.GetIsFollow(context.Background(), &relation.RelationIsFollowRequest{
		MyId:   myId,
		UserId: userId,
	})
	userInfo.IsFollow = resp3.GetIsFollow()
	// TODO 作品数，喜欢数
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
	return "https://th.bing.com/th/id/R.e6b23f7279370871e1d13a9b8472bacc?rik=8%2fOfHw3gtBb%2fiw&riu=http%3a%2f%2fi2.hdslb.com%2fbfs%2farchive%2f63cd640f4ba78525a8797a94888d0fac654d7cdb.jpg&ehk=td%2bSb3B1RLfBQKrbTv43cN3r7MmjFMxg07MvGUDVoao%3d&risl=&pid=ImgRaw&r=0"
}
