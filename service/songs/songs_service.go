package songs

import (
	"context"
	"login-app/model/web/songs"
)

type SongsService interface {
	Create(ctx context.Context, request songs.SongCreateRequest) songs.SongResponse
	Update(ctx context.Context, request songs.SongUpdateRequest) songs.SongResponse
	Delete(ctx context.Context, songId int)
	FindById(ctx context.Context, songId int) songs.SongResponse
	FindAll(ctx context.Context) []songs.SongResponse
}
