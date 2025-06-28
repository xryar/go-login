package web

type AlbumUpdateRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Year string `json:"year"`
}
