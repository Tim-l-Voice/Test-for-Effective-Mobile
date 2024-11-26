package main

import (
	"database/sql"
	"log"
	_ "song_library/docs"
	"song_library/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db *sql.DB

// @title Music Library API
// @version 1.0
// @description API для управления музыкальной библиотекой
// @BasePath /
// @host localhost:8080
func main() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=2090 dbname=postgres sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Successfully connected to the database")

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/songs", handlers.GetSongs)
	r.POST("/songs", handlers.AddSong)

	createTableIfNotExists()

	r.Run(":8080")
}

func createTableIfNotExists() {
	query := `
	CREATE TABLE IF NOT EXISTS songs (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		artist VARCHAR(255) NOT NULL
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	log.Println("Table 'songs' is ready or created.")
}
