package app

import (
	"login-app/controller"
	"login-app/exception"
	"login-app/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(usersController controller.UsersController, albumsController controller.AlbumsController) *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/register", usersController.Create)
	router.POST("/api/login", usersController.Login)

	albumsRouter(router, albumsController)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func albumsRouter(router *httprouter.Router, albumsController controller.AlbumsController) {
	albumRouter := httprouter.New()
	albumRouter.POST("/api/albums", albumsController.Create)
	albumRouter.GET("/api/albums", albumsController.FindAll)
	albumRouter.GET("/api/albums/:albumId", albumsController.FindById)
	albumRouter.PUT("/api/albums/:albumId", albumsController.Update)
	albumRouter.DELETE("/api/albums/:albumId", albumsController.Delete)

	protectedHandler := middleware.NewAuthMiddleware(albumRouter)
	router.Handler("POST", "/api/albums", protectedHandler)
	router.Handler("GET", "/api/albums", protectedHandler)
	router.Handler("GET", "/api/albums/:albumId", protectedHandler)
	router.Handler("PUT", "/api/albums/:albumId", protectedHandler)
	router.Handler("DELETE", "/api/albums/:albumId", protectedHandler)
}
