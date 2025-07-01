package main

import (
	"fmt"
	"login-app/app"
	albumsController "login-app/controller/albums"
	songsController "login-app/controller/songs"
	usersController "login-app/controller/users"
	"login-app/helper"
	albumsRepository "login-app/repository/albums"
	songsRepository "login-app/repository/songs"
	usersRepository "login-app/repository/users"
	albumsService "login-app/service/albums"
	songsService "login-app/service/songs"
	usersService "login-app/service/users"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	usersRepository := usersRepository.NewUsersRepository()
	albumsRepository := albumsRepository.NewAlbumsRepository()
	songsRepository := songsRepository.NewSongsRepository()
	usersService := usersService.NewUsersService(usersRepository, db, validate)
	albumsService := albumsService.NewAlbumsService(albumsRepository, db, validate)
	songsService := songsService.NewSongsService(songsRepository, db, validate)
	usersController := usersController.NewUsersController(usersService)
	albumsController := albumsController.NewAlbumController(albumsService)
	songsController := songsController.NewSongsController(songsService)
	router := app.NewRouter(usersController, albumsController, songsController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("starting web server at http://localhost:3000/")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
