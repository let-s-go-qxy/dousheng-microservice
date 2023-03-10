package main

import (
	"context"
	"dousheng/cmd/user/internal/service"
	"dousheng/kitex_gen/user"
	g "dousheng/pkg/global"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	resp = &user.UserInfoResponse{}
	user1 := user.User{}
	user1, err = service.UserInfo(req.GetMyId(), req.GetUserId())
	if err != nil {
		return nil, err
	}
	resp.User = &user1
	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp = &user.UserLoginResponse{}
	resp.UserId, resp.Token, err = service.UserLogin(req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	resp.StatusCode = g.StatusOk
	resp.StatusMsg = "ok"
	return
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	resp = &user.UserRegisterResponse{}
	resp.UserId, resp.Token, err = service.UserRegister(req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	resp.StatusCode = g.StatusOk
	resp.StatusMsg = "ok"
	return
}

// GetAvatar implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetAvatar(ctx context.Context, req *user.UserAvatarRequest) (resp *user.UserAvatarResponse, err error) {
	resp = &user.UserAvatarResponse{}
	resp.Avatar = service.GetAvatar(req.UserId)
	return
}

// GetBackgroundImage implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetBackgroundImage(ctx context.Context, req *user.UserBackgroundImageRequest) (resp *user.UserBackgroundImageResponse, err error) {
	resp = &user.UserBackgroundImageResponse{}
	resp.BackgroundImage = service.GetBackgroundImage(req.UserId)
	return resp, err
}
