package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UsersController interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
