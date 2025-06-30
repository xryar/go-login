package service

import (
	"context"
	"login-app/model/web"
)

type UsersService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Login(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse
}
