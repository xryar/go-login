package main

import (
	"fmt"
	"login-app/app"
	"login-app/controller"
	"login-app/helper"
	"login-app/repository"
	"login-app/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	usersRepository := repository.NewUsersRepository()
	albumsRepository := repository.NewAlbumsRepository()
	usersService := service.NewUsersService(usersRepository, db, validate)
	AlbumsService := service.NewAlbumsService(albumsRepository, db, validate)
	usersController := controller.NewUsersController(usersService)
	albumsController := controller.NewAlbumController(AlbumsService)
	router := app.NewRouter(usersController, albumsController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("starting web server at http://localhost:3000/")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
