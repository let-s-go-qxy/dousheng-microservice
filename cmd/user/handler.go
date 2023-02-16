package main

import (
	"context"
	user "dousheng/kitex_gen/user"
	"github.com/pkg/errors"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	resp = &user.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		User: &user.User{
			Id:            1,
			Name:          "aei",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		},
	}
	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	if req.GetUsername() == "" || req.GetPassword() == "" {
		err = errors.New("账号或密码不规范")
	}
	return
}
