package main

import (
	"fmt"
	"login-app/app"
	albumsController "login-app/controller/albums"
	usersController "login-app/controller/users"
	"login-app/helper"
	albumsRepository "login-app/repository/albums"
	usersRepository "login-app/repository/users"
	albumsService "login-app/service/albums"
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
	usersService := usersService.NewUsersService(usersRepository, db, validate)
	AlbumsService := albumsService.NewAlbumsService(albumsRepository, db, validate)
	usersController := usersController.NewUsersController(usersService)
	albumsController := albumsController.NewAlbumController(AlbumsService)
	router := app.NewRouter(usersController, albumsController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("starting web server at http://localhost:3000/")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
