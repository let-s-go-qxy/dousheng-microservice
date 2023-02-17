package service

import (
	"dousheng/cmd/user/internal/model"
	"dousheng/conf"
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
	user := &User{
		Name: name,
	}
	err = GetUser(user)
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
	userId = int64(user.Id)
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

func UserInfo(myId int64, userId int64) (Id int64, FollowCount, FollowerCount int32, Name string, IsFollow bool, avatar string, err error) {
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
	avatar = GetAvatar(userId)
	Name = userDao.Name
	//FollowCount = int(model.GetFollowCount(userDao.Id))
	//FollowerCount = int(model.GetFollowerCount(userDao.Id))
	//IsFollow = model.IsFollow(myId, userId)
	return userDao.Id, FollowCount, FollowerCount, Name, IsFollow, avatar, err
}

// GetAvatar 获取用户头像
func GetAvatar(userID int64) string {
	strUserID := strconv.Itoa(int(userID))
	avatarURL := conf.OSSAvatarPreURL + strUserID + "_avatar.jpg"
	return avatarURL
}
