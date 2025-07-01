package app

import (
	albumsController "login-app/controller/albums"
	songsController "login-app/controller/songs"
	usersController "login-app/controller/users"
	"login-app/exception"
	"login-app/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(usersController usersController.UsersController, albumsController albumsController.AlbumsController, songsController songsController.SongsController) *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/register", usersController.Create)
	router.POST("/api/login", usersController.Login)

	albumsRouter(router, albumsController)
	songsRouter(router, songsController)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func albumsRouter(router *httprouter.Router, albumsController albumsController.AlbumsController) {
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

func songsRouter(router *httprouter.Router, songsController songsController.SongsController) {
	songRouter := httprouter.New()
	songRouter.POST("/api/songs", songsController.Create)
	songRouter.GET("/api/songs", songsController.FindAll)
	songRouter.GET("/api/songs/:songId", songsController.FindById)
	songRouter.PUT("/api/songs/:songId", songsController.Update)
	songRouter.DELETE("/api/songs/:songId", songsController.Delete)

	protectedHandler := middleware.NewAuthMiddleware(songRouter)
	router.Handler("POST", "/api/songs", protectedHandler)
	router.Handler("GET", "/api/songs", protectedHandler)
	router.Handler("GET", "/api//:songId", protectedHandler)
	router.Handler("PUT", "/api//:songId", protectedHandler)
	router.Handler("DELETE", "/api//:songId", protectedHandler)

}
