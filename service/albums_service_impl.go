package service

import (
	"context"
	"database/sql"
	"login-app/exception"
	"login-app/helper"
	"login-app/model/domain"
	"login-app/model/web"
	"login-app/repository"

	"github.com/go-playground/validator/v10"
)

type AlbumServiceImpl struct {
	AlbumRepository repository.AlbumsRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewAlbumService(albumRepository repository.AlbumsRepository, DB *sql.DB, validate *validator.Validate) *AlbumServiceImpl {
	return &AlbumServiceImpl{
		AlbumRepository: albumRepository,
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

	album := domain.Albums{
		Name: request.Name,
		Year: request.Year,
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

func (service *AlbumServiceImpl) FindById(ctx context.Context, albumId int) web.AlbumResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	album, err := service.AlbumRepository.FindById(ctx, tx, albumId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAlbumResponse(album)
}

func (service *AlbumServiceImpl) FindAll(ctx context.Context) []web.AlbumResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	albums := service.AlbumRepository.FindAll(ctx, tx)

	return helper.ToAlbumResponses(albums)
}
