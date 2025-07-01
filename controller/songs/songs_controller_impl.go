package songs

import (
	"login-app/helper"
	"login-app/model/web"
	"login-app/model/web/songs"
	service "login-app/service/songs"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type SongsControllerImpl struct {
	SongsService service.SongsService
}

func NewSongsController(songsService service.SongsService) *SongsControllerImpl {
	return &SongsControllerImpl{
		SongsService: songsService,
	}
}

func (controller *SongsControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	songCreateRequest := songs.SongCreateRequest{}
	helper.ReadFromRequestBody(r, &songCreateRequest)

	songResponse := controller.SongsService.Create(r.Context(), songCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   songResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *SongsControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	songUpdateRequest := songs.SongUpdateRequest{}
	helper.ReadFromRequestBody(r, &songUpdateRequest)

	songId := p.ByName("songId")
	id, err := strconv.Atoi(songId)
	helper.PanicIfError(err)

	songUpdateRequest.Id = id

	songResponse := controller.SongsService.Update(r.Context(), songUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   songResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *SongsControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	songId := p.ByName("songId")
	id, err := strconv.Atoi(songId)
	helper.PanicIfError(err)

	controller.SongsService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *SongsControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	songId := p.ByName("songId")
	id, err := strconv.Atoi(songId)
	helper.PanicIfError(err)

	songResponse := controller.SongsService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   songResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *SongsControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	songResponse := controller.SongsService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   songResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
