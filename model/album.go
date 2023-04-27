package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

var dbpool *pgxpool.Pool

func ConnectToDB(database_url string) error {
	var err error
	dbpool, err = pgxpool.New(context.Background(), database_url)
	if err != nil {
		return err
	}
	return nil
}

func AllAlbums() ([]Album, error) {
	var albums []Album

	rows, _ := dbpool.Query(context.Background(), "SELECT * FROM album")
	albums, err := pgx.CollectRows(rows, pgx.RowToStructByName[Album])
	if err != nil {
		return nil, fmt.Errorf("getAllAlbums %q", err)
	}
	defer rows.Close()

	return albums, err
}

func AlbumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows
	var albums []Album

	rows, _ := dbpool.Query(context.Background(), "SELECT * FROM album WHERE artist=$1", name)
	albums, err := pgx.CollectRows(rows, pgx.RowToStructByName[Album])
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	return albums, nil
}

// Queries for the album with a specified ID
func AlbumsByID(id int64) (Album, error) {

	// an album to hold data from the returned row
	var alb Album

	row := dbpool.QueryRow(context.Background(), "SELECT * FROM album WHERE id = $1", id)

	if err := row.Scan(&alb.ID, &alb.Artist, &alb.Title, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsByID %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d, %v", id, err)
	}

	return alb, nil
}

// addAlbum adds the specified album to the database
// returning the album ID of the new entry
func AddAlbum(alb Album) (int64, error) {

	result, err := dbpool.Exec(context.Background(), "INSERT INTO album (title, artist, price) VALUES ($1, $2, $3)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	fmt.Printf("Result was: %v", result)
	return 1, nil
}
