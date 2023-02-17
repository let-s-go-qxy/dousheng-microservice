package service

import (
	"dousheng/cmd/user/internal/model"
	"dousheng/conf"
	userService "dousheng/kitex_gen/user"
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

func UserInfo(myId int64, userId int64) (userInfo userService.User, err error) {
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
	userInfo = userService.User{
		Id:              userId,
		Name:            userDao.Name,
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		WorkCount:       0,
		BackgroundImage: "https://tse2-mm.cn.bing.net/th/id/OIP-C.BRuY39z2iJY_hiqkoNhH_wHaE7?pid=ImgDet&rs=1",
		Signature:       "我不想再当一个xx了，我只想拥有快乐（谢谢你，狗子）",
		TotalFavorite:   0,
		FavoriteCount:   0,
		Avatar:          "",
	}
	userInfo.Avatar = GetAvatar(userId)
	userInfo.Name = userDao.Name
	//FollowCount = int(model.GetFollowCount(userDao.Id))
	//FollowerCount = int(model.GetFollowerCount(userDao.Id))
	//IsFollow = model.IsFollow(myId, userId)
	return
}

// GetAvatar 获取用户头像
func GetAvatar(userID int64) string {
	strUserID := strconv.Itoa(int(userID))
	avatarURL := conf.OSSAvatarPreURL + strUserID + "_avatar.jpg"
	return avatarURL
}
