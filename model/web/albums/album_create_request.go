package web

type AlbumCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
	Year string `validate:"required,min=1" json:"year"`
}
