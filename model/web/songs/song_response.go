package songs

type SongResponse struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Year      string `json:"year"`
	Genre     string `json:"genre"`
	Performer string `json:"performer"`
	Duration  int    `json:"duration"`
	AlbumId   int    `json:"album_id"`
}
