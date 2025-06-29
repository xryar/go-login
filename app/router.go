package app

import (
	"login-app/controller"
	"login-app/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(usersController controller.UsersController, albumsController controller.AlbumsController) *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/register", usersController.Create)
	router.POST("/api/login", usersController.Login)

	router.POST("/api/albums", albumsController.Create)
	router.GET("/api/albums", albumsController.FindAll)
	router.GET("/api/albums/:albumId", albumsController.FindById)
	router.PUT("/api/albums/:albumId", albumsController.Update)
	router.DELETE("/api/albums/:albumId", albumsController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
