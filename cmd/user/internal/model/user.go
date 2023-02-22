package model

import (
	g "dousheng/pkg/global"
	"gorm.io/gorm"
)

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

// CreateUser 增加用户
func CreateUser(user *User) (db *gorm.DB, err error) {
	var count int64
	db = g.WriteMysqlDB.Model(&User{}).Where("name = ?", user.Name).Count(&count)
	if count > 0 {
		err = g.ErrDbCreateUniqueKeyRepeatedly
		return
	}
	db = g.WriteMysqlDB.Create(user)
	err = db.Error
	return
}

// GetUser 通过名称和user_id查询记录 limit 1
func GetUser(user *User) (err error) {
	if user.Name != "" {
		err = g.ReadMysqlDB.First(user, "name = ?", user.Name).Error
		return
	}
	err = g.ReadMysqlDB.First(user, "id = ?", user.Id).Error
	return
}

func GetNameById(userId int) string {
	var user User
	g.ReadMysqlDB.Where("id = ?", userId).Take(&user)
	return user.Name
}
