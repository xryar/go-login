package service

import (
	"context"
	"database/sql"
	"login-app/exception"
	"login-app/helper"
	"login-app/model/domain"
	web "login-app/model/web/albums"
	repository "login-app/repository/albums"
	songRepository "login-app/repository/songs"

	"github.com/go-playground/validator/v10"
)

type AlbumServiceImpl struct {
	AlbumRepository repository.AlbumsRepository
	SongRepository  songRepository.SongsRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewAlbumsService(albumRepository repository.AlbumsRepository, songRepository songRepository.SongsRepository, DB *sql.DB, validate *validator.Validate) *AlbumServiceImpl {
	return &AlbumServiceImpl{
		AlbumRepository: albumRepository,
		SongRepository:  songRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *AlbumServiceImpl) Create(ctx context.Context, request web.AlbumCreateRequest) web.AlbumResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userId := helper.GetUserIdFromContext(ctx)

	album := domain.Albums{
		Name:   request.Name,
		Year:   request.Year,
		UserId: userId,
	}

	album = service.AlbumRepository.Save(ctx, tx, album)

	return helper.ToAlbumResponse(album)
}

func (service *AlbumServiceImpl) Update(ctx context.Context, request web.AlbumUpdateRequest) web.AlbumResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	album, err := service.AlbumRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	album.Name = request.Name
	album.Year = request.Year

	album = service.AlbumRepository.Update(ctx, tx, album)

	return helper.ToAlbumResponse(album)
}

func (service *AlbumServiceImpl) Delete(ctx context.Context, albumId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	album, err := service.AlbumRepository.FindById(ctx, tx, albumId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.AlbumRepository.Delete(ctx, tx, album)
}

func (service *AlbumServiceImpl) FindById(ctx context.Context, albumId int) web.AlbumWithSongResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	album, err := service.AlbumRepository.FindById(ctx, tx, albumId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	songs := service.SongRepository.FindByAlbumId(ctx, tx, albumId)

	var songResponses []web.SongInAlbumResponse
	for _, song := range songs {
		songResponses = append(songResponses, web.SongInAlbumResponse{
			Id:        song.Id,
			Title:     song.Title,
			Performer: song.Performer,
		})
	}

	return web.AlbumWithSongResponse{
		Id:    album.Id,
		Name:  album.Name,
		Year:  album.Year,
		Songs: songResponses,
	}
}

func (service *AlbumServiceImpl) FindAll(ctx context.Context) []web.AlbumResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	albums := service.AlbumRepository.FindAll(ctx, tx)

	return helper.ToAlbumResponses(albums)
}
