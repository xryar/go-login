package repository

import (
	"context"
	"database/sql"
	"login-app/helper"
	"login-app/model/domain"
)

type UsersRepositoryImpl struct {
}

func NewUsersRepository() *UsersRepositoryImpl {
	return &UsersRepositoryImpl{}
}

func (repository *UsersRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users {
	SQL := "INSERT INTO users(username, fullname, password) values(?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Fullname, user.Password)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}
