package songs

import (
	"login-app/helper"
	"login-app/model/web"
	"login-app/model/web/songs"
	service "login-app/service/songs"
	"net/http"

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
