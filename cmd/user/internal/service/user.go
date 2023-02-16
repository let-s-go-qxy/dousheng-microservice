package service

import (
	"dousheng/cmd/user/internal/model"
	utils2 "dousheng/pkg/utils"
	"errors"
	"gorm.io/gorm"
)

// JudgeNameAndPassword 判断用户名和密码是否符合要求
func JudgeNameAndPassword(name, password string) bool {
	if len(name) > 32 || len(password) > 32 || name == "" || len(password) < 6 {
		return false
	}
	return true
}

func UserLogin(name, password string) (userId int, token string, err error) {
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
