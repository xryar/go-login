package main

import (
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
	usersService := service.NewUsersService(usersRepository, db, validate)
	usersController := controller.NewUsersController(usersService)
	router := app.NewRouter(usersController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
