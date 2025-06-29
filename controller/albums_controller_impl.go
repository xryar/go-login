package controller

import (
	"login-app/helper"
	"login-app/model/web"
	"login-app/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type AlbumControllerImpl struct {
	AlbumService service.AlbumsService
}

func NewAlbumController(albumService service.AlbumsService) *AlbumControllerImpl {
	return &AlbumControllerImpl{
		AlbumService: albumService,
	}
}

func (controller *AlbumControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	albumCreateRequest := web.AlbumCreateRequest{}
	helper.ReadFromRequestBody(r, &albumCreateRequest)

	albumResponse := controller.AlbumService.Create(r.Context(), albumCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   albumResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *AlbumControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	albumUpdateRequest := web.AlbumUpdateRequest{}
	helper.ReadFromRequestBody(r, &albumUpdateRequest)

	albumId := p.ByName("albumId")
	id, err := strconv.Atoi(albumId)
	helper.PanicIfError(err)

	albumUpdateRequest.Id = id

	albumResponse := controller.AlbumService.Update(r.Context(), albumUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   albumResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (controller *AlbumControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	albumId := p.ByName("albumId")
	id, err := strconv.Atoi(albumId)
	helper.PanicIfError(err)

	controller.AlbumService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *AlbumControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	albumId := p.ByName("albumId")
	id, err := strconv.Atoi(albumId)
	helper.PanicIfError(err)

	albumResponse := controller.AlbumService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   albumResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *AlbumControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	albumResponse := controller.AlbumService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   albumResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
