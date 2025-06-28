package app

import (
	"login-app/controller"
	"login-app/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(usersController controller.UsersController) *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/register", usersController.Create)
	router.POST("/api/login", usersController.Login)

	router.PanicHandler = exception.ErrorHandler

	return router
}
