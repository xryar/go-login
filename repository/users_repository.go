package repository

import (
	"context"
	"database/sql"
	"login-app/model/domain"
)

type UsersRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.Users, error)
}
