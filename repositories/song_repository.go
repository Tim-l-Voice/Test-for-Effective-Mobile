package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"song_library/models"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	connStr := "user=postgres password=2090 dbname=song_library sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
		return err
	}

	log.Println("Successfully connected to the database")
	return nil
}

func GetAllSongs(page string, limit string) ([]models.Song, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, fmt.Errorf("invalid page parameter: %v", err)
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return nil, fmt.Errorf("invalid limit parameter: %v", err)
	}

	offset := (pageInt - 1) * limitInt

	rows, err := db.Query("SELECT id, title, artist FROM songs LIMIT $1 OFFSET $2", limitInt, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.ID, &song.Title, &song.Artist); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func AddSong(song models.Song) error {
	_, err := db.Exec("INSERT INTO songs (title, artist) VALUES ($1, $2)", song.Title, song.Artist)
	return err
}
