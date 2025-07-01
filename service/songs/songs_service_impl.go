package songs

import (
	"context"
	"database/sql"
	"login-app/model/web/songs"
	repository "login-app/repository/songs"

	"github.com/go-playground/validator/v10"
)

type SongsServiceImpl struct {
	SongsRepository repository.SongsRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewSongsService(songsRepository repository.SongsRepository, DB *sql.DB, validate *validator.Validate) *SongsServiceImpl {
	return &SongsServiceImpl{
		SongsRepository: songsRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *SongsServiceImpl) Create(ctx context.Context, request songs.SongCreateRequest) songs.SongResponse {

}

func (service *SongsServiceImpl) Update(ctx context.Context, request songs.SongUpdateRequest) songs.SongResponse {

}

func (service *SongsServiceImpl) Delete(ctx context.Context, songId int) {

}

func (service *SongsServiceImpl) FindById(ctx context.Context, songId int) songs.SongResponse {

}

func (service *SongsServiceImpl) FindAll(ctx context.Context) []songs.SongResponse {

}
