package service

import (
	"context"
	web "login-app/model/web/users"
)

type UsersService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Login(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse
}
