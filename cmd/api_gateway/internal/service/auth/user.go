package auth

import (
	g "dousheng/pkg/global"
	"errors"
	"github.com/form3tech-oss/jwt-go"
)

type UserClaims struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Expire int    `json:"expire"`
	jwt.StandardClaims
}

// ParseToken 解析token
func ParseToken(str string) (UserClaims, error) {
	c := new(UserClaims)
	token, err := jwt.ParseWithClaims(str, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(g.Config.JwtSecretKey), nil
	})
	if err != nil {
		return *c, errors.New("token不合法")
	}
	if token.Valid != true {
		return *c, errors.New("token不合法")
	}
	return *c, err
}
