package songs

import (
	"context"
	"database/sql"
	"login-app/exception"
	"login-app/helper"
	"login-app/model/domain"
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
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userId := helper.GetUserIdFromContext(ctx)

	song := domain.Songs{
		Title:     request.Title,
		Year:      request.Year,
		Genre:     request.Genre,
		Performer: request.Performer,
		Duration:  request.Duration,
		AlbumId:   request.AlbumId,
		UserId:    userId,
	}

	song = service.SongsRepository.Save(ctx, tx, song)

	return helper.ToSongResponse(song)
}

func (service *SongsServiceImpl) Update(ctx context.Context, request songs.SongUpdateRequest) songs.SongResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	song, err := service.SongsRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	song.Title = request.Title
	song.Year = request.Year
	song.Genre = request.Genre
	song.Performer = request.Performer
	song.Duration = request.Duration
	song.AlbumId = request.AlbumId

	song = service.SongsRepository.Update(ctx, tx, song)

	return helper.ToSongResponse(song)
}

func (service *SongsServiceImpl) Delete(ctx context.Context, songId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	song, err := service.SongsRepository.FindById(ctx, tx, songId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.SongsRepository.Delete(ctx, tx, song)

}

func (service *SongsServiceImpl) FindById(ctx context.Context, songId int) songs.SongResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	song, err := service.SongsRepository.FindById(ctx, tx, songId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToSongResponse(song)
}

func (service *SongsServiceImpl) FindAll(ctx context.Context) []songs.SongResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	songs := service.SongsRepository.FindAll(ctx, tx)

	return helper.ToSongResponses(songs)
}
