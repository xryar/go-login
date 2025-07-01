package songs

import service "login-app/service/songs"

type SongsControllerImpl struct {
	SongsService service.SongsService
}

func NewSongsController(songsService service.SongsService) *SongsControllerImpl {
	return &SongsControllerImpl{
		SongsService: songsService,
	}
}
