package helper

import (
	"login-app/model/domain"
	albumResponse "login-app/model/web/albums"
	songResponse "login-app/model/web/songs"
	userResponse "login-app/model/web/users"
)

func ToUserResponse(user domain.Users) userResponse.UserResponse {
	return userResponse.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Fullname: user.Fullname,
	}
}

func ToUserLoginResponse(token string) userResponse.UserLoginResponse {
	return userResponse.UserLoginResponse{
		Token: token,
	}
}

func ToAlbumResponse(album domain.Albums) albumResponse.AlbumResponse {
	return albumResponse.AlbumResponse{
		Id:   album.Id,
		Name: album.Name,
		Year: album.Year,
	}
}

func ToAlbumResponses(albums []domain.Albums) []albumResponse.AlbumResponse {
	var albumResponses []albumResponse.AlbumResponse
	for _, album := range albums {
		albumResponses = append(albumResponses, ToAlbumResponse(album))
	}

	return albumResponses
}

func ToAlbumWithSongResponse(album domain.Albums, songs []domain.Songs) albumResponse.AlbumWithSongResponse {
	var songResponses []albumResponse.SongInAlbumResponse
	for _, song := range songs {
		songResponses = append(songResponses, albumResponse.SongInAlbumResponse{
			Id:        song.Id,
			Title:     song.Title,
			Performer: song.Performer,
		})
	}

	return albumResponse.AlbumWithSongResponse{
		Id:    album.Id,
		Name:  album.Name,
		Year:  album.Year,
		Songs: songResponses,
	}
}

func ToSongResponse(song domain.Songs) songResponse.SongResponse {
	return songResponse.SongResponse{
		Id:        song.Id,
		Title:     song.Title,
		Year:      song.Year,
		Genre:     song.Genre,
		Performer: song.Performer,
		Duration:  song.Duration,
		AlbumId:   song.AlbumId,
	}
}

func ToSongResponses(songs []domain.Songs) []songResponse.SongResponse {
	var songResponses []songResponse.SongResponse
	for _, song := range songs {
		songResponses = append(songResponses, ToSongResponse(song))
	}

	return songResponses
}
