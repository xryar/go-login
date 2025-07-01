package songs

type SongCreateRequest struct {
	Title     string `validate:"required,min=1,max=100" json:"title"`
	Year      string `validate:"required,min=1" json:"year"`
	Genre     string `validate:"required,min=1" json:"genre"`
	Performer string `validate:"required,min=1" json:"performer"`
	Duration  int    `validate:"required,min=1" json:"duration"`
	AlbumId   int    `validate:"required,min=1" json:"album_id"`
}
