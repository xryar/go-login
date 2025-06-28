package helper

import (
	"login-app/model/domain"
	"login-app/model/web"
)

func ToUserResponse(user domain.Users) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Fullname: user.Fullname,
	}
}

func ToUserLoginResponse(token string) web.UserLoginResponse {
	return web.UserLoginResponse{
		Token: token,
	}
}
