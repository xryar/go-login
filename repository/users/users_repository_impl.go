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
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *UsersRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.Users, error) {
	SQL := "SELECT * FROM users WHERE username = (?) LIMIT 1"
	row := tx.QueryRowContext(ctx, SQL, username)

	user := domain.Users{}
	err := row.Scan(&user.Id, &user.Username, &user.Fullname, &user.Password)
	helper.PanicIfError(err)

	return user, nil
}
