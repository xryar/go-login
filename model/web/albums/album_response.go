package web

type SongInAlbumResponse struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Performer string `json:"performer"`
}

type AlbumResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Year string `json:"year"`
}

type AlbumWithSongResponse struct {
	Id    int                   `json:"id"`
	Name  string                `json:"name"`
	Year  string                `json:"year"`
	Songs []SongInAlbumResponse `json:"songs"`
}
