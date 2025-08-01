package songs

import (
	"context"
	"database/sql"
	"errors"
	"login-app/helper"
	"login-app/model/domain"
)

type SongsRepositoryImpl struct {
}

func NewSongsRepository() *SongsRepositoryImpl {
	return &SongsRepositoryImpl{}
}

func (repository *SongsRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, songs domain.Songs) domain.Songs {
	SQL := "INSERT INTO songs(title, year, genre, performer, duration, album_id, user_id) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, songs.Title, songs.Year, songs.Genre, songs.Performer, songs.Duration, songs.AlbumId, songs.UserId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	songs.Id = int(id)
	return songs
}

func (repository *SongsRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, songs domain.Songs) domain.Songs {
	SQL := "UPDATE songs set title = ?, year = ?, genre = ?, performer = ?, duration = ?, album_id = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, songs.Title, songs.Year, songs.Genre, songs.Performer, songs.Duration, songs.AlbumId, songs.Id)
	helper.PanicIfError(err)

	return songs
}

func (repository *SongsRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, songs domain.Songs) {
	SQL := "DELETE FROM songs WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, songs.Id)
	helper.PanicIfError(err)
}

func (repository *SongsRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, songId int) (domain.Songs, error) {
	SQL := "SELECT id, title, year, genre, performer, duration, album_id FROM songs WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, songId)
	helper.PanicIfError(err)
	defer rows.Close()

	song := domain.Songs{}
	if rows.Next() {
		rows.Scan(&song.Id, &song.Title, &song.Year, &song.Genre, &song.Performer, &song.Duration, &song.AlbumId)
		helper.PanicIfError(err)
		return song, nil
	} else {
		return song, errors.New("song is not found")
	}
}

func (repository *SongsRepositoryImpl) FindByAlbumId(ctx context.Context, tx *sql.Tx, albumId int) []domain.Songs {
	SQL := "SELECT id, title, performer FROM songs WHERE album_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, albumId)
	helper.PanicIfError(err)
	defer rows.Close()

	var songs []domain.Songs
	for rows.Next() {
		song := domain.Songs{}
		err := rows.Scan(&song.Id, &song.Title, &song.Performer)
		helper.PanicIfError(err)
		songs = append(songs, song)
	}

	return songs
}

func (repository *SongsRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Songs {
	SQL := "SELECT id, title, year, genre, performer, duration, album_id FROM songs"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var songs []domain.Songs
	for rows.Next() {
		song := domain.Songs{}
		err := rows.Scan(&song.Id, &song.Title, &song.Year, &song.Genre, &song.Performer, &song.Duration, &song.AlbumId)
		helper.PanicIfError(err)
		songs = append(songs, song)
	}

	return songs
}
