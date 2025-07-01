package songs

import (
	"context"
	"database/sql"
	"login-app/model/domain"
)

type SongsRepository interface {
	Save(ctx context.Context, tx *sql.Tx, songs domain.Songs) domain.Songs
	Update(ctx context.Context, tx *sql.Tx, songs domain.Songs) domain.Songs
	Delete(ctx context.Context, tx *sql.Tx, songs domain.Songs)
	FindById(ctx context.Context, tx *sql.Tx, songId int) (domain.Songs, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Songs
}
