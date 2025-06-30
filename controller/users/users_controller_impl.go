package controller

import (
	"login-app/helper"
	"login-app/model/web"
	service "login-app/service/users"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UsersControllerImpl struct {
	UsersService service.UsersService
}

func NewUsersController(usersService service.UsersService) *UsersControllerImpl {
	return &UsersControllerImpl{
		UsersService: usersService,
	}
}

func (controller *UsersControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(r, &userCreateRequest)

	userResponse := controller.UsersService.Create(r.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UsersControllerImpl) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userLoginRequest := web.UserLoginRequest{}
	helper.ReadFromRequestBody(r, &userLoginRequest)

	userResponse := controller.UsersService.Login(r.Context(), userLoginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "SUCCESS",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
