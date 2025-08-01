package service

import (
	"context"
	web "login-app/model/web/albums"
)

type AlbumsService interface {
	Create(ctx context.Context, request web.AlbumCreateRequest) web.AlbumResponse
	Update(ctx context.Context, request web.AlbumUpdateRequest) web.AlbumResponse
	Delete(ctx context.Context, albumId int)
	FindById(ctx context.Context, albumId int) web.AlbumWithSongResponse
	FindAll(ctx context.Context) []web.AlbumResponse
}
