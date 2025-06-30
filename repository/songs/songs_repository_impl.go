package songs

import (
	"context"
	"database/sql"
	"login-app/model/domain"
)

type SongsRepositoryImpl struct {
}

func NewSongsRepository() *SongsRepositoryImpl {
	return &SongsRepositoryImpl{}
}

func (repository *SongsRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, songs domain.Songs) domain.Songs {

}

func (repository *SongsRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, songs domain.Songs) domain.Songs {

}

func (repository *SongsRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, songs domain.Songs) {

}

func (repository *SongsRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, songId int) (domain.Songs, error) {

}

func (repository *SongsRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Songs {

}
