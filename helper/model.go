package helper

import (
	"login-app/model/domain"
	"login-app/model/web"
)

func ToUserResponse(user domain.Users) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Fullname: user.Fullname,
	}
}

func ToUserLoginResponse(token string) web.UserLoginResponse {
	return web.UserLoginResponse{
		Token: token,
	}
}

func ToAlbumResponse(album domain.Albums) web.AlbumResponse {
	return web.AlbumResponse{
		Id:   album.Id,
		Name: album.Name,
		Year: album.Year,
	}
}

func ToAlbumResponses(albums []domain.Albums) []web.AlbumResponse {
	var albumResponses []web.AlbumResponse
	for _, album := range albums {
		albumResponses = append(albumResponses, ToAlbumResponse(album))
	}

	return albumResponses
}
