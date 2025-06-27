package service

import (
	"context"
	"database/sql"
	"login-app/helper"
	"login-app/model/domain"
	"login-app/model/web"
	"login-app/repository"

	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewUsersService(usersRepository repository.UsersRepository, DB *sql.DB, validate *validator.Validate) *UsersServiceImpl {
	return &UsersServiceImpl{
		UsersRepository: usersRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *UsersServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.Users{
		Username: request.Username,
		Fullname: request.Fullname,
		Password: request.Password,
	}

	user = service.UsersRepository.Create(ctx, tx, user)

	return helper.ToUserResponse(user)
}
