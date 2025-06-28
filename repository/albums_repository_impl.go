package repository

import (
	"context"
	"database/sql"
	"errors"
	"login-app/helper"
	"login-app/model/domain"
)

type AlbumsRepositoryImpl struct {
}

func NewAlbumsRepository() *AlbumsRepositoryImpl {
	return &AlbumsRepositoryImpl{}
}

func (repository *AlbumsRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, albums domain.Albums) domain.Albums {
	SQL := "INSERT INTO albums(name, year) values (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, albums.Name, albums.Year)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	albums.Id = int(id)
	return albums
}

func (repository *AlbumsRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, albums domain.Albums) domain.Albums {
	SQL := "UPDATE albums SET name = ? AND year = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, albums.Name, albums.Year, albums.Id)
	helper.PanicIfError(err)

	return albums
}

func (repository *AlbumsRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, albums domain.Albums) {
	SQL := "DELETE FROM albums WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, albums.Id)
	helper.PanicIfError(err)
}

func (repository *AlbumsRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, albumId int) (domain.Albums, error) {
	SQL := "SELECT * FROM albums WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, albumId)
	helper.PanicIfError(err)
	defer rows.Close()

	album := domain.Albums{}
	if rows.Next() {
		rows.Scan(&album.Id, &album.Name, &album.Year)
		helper.PanicIfError(err)
		return album, nil
	} else {
		return album, errors.New("album is not found")
	}
}

func (repository *AlbumsRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Albums {
	SQL := "SELECT * FROM albums"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var albums []domain.Albums
	for rows.Next() {
		album := domain.Albums{}
		err := rows.Scan(&album.Id, &album.Name, &album.Year)
		helper.PanicIfError(err)
		albums = append(albums, album)
	}

	return albums
}
