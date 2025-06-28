package repository

import (
	"context"
	"database/sql"
	"login-app/model/domain"
)

type AlbumsRepository interface {
	Save(ctx context.Context, tx *sql.Tx, albums domain.Albums) domain.Albums
	Update(ctx context.Context, tx *sql.Tx, albums domain.Albums) domain.Albums
	Delete(ctx context.Context, tx *sql.Tx, albums domain.Albums)
	FindById(ctx context.Context, tx *sql.Tx, albums domain.Albums) (domain.Albums, error)
	FindAll(ctx context.Context, tx *sql.Tx, albums domain.Albums) []domain.Albums
}
