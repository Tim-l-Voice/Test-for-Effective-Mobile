package handlers

import (
	"net/http"
	"song_library/models"
	"song_library/repositories"
	"song_library/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Get all songs
// @Description Fetch all the songs from the music library
// @Tags songs
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Limit of songs per page" default(10)
// @Success 200 {array} models.Song
// @Failure 500 {object} utils.ErrorResponse
// @Router /songs [get]
func GetSongs(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	songs, err := repositories.GetAllSongs(page, limit)
	if err != nil {
		utils.Logger.Errorf("Error fetching songs: %v", err)
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse{Message: "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, songs)
}

// @Summary Add a new song
// @Description Add a new song to the music library
// @Tags songs
// @Accept json
// @Produce json
// @Param song body models.SongRequest true "Song details"
// @Success 201 {object} models.Song
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /songs [post]
func AddSong(c *gin.Context) {
	var newSong models.SongRequest
	if err := c.ShouldBindJSON(&newSong); err != nil {
		utils.Logger.Errorf("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Message: "Invalid request body"})
		return
	}

	song := models.Song{
		Title:  newSong.Title,
		Artist: newSong.Artist,
	}

	err := repositories.AddSong(song)
	if err != nil {
		utils.Logger.Errorf("Error adding song: %v", err)
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse{Message: "Failed to add song"})
		return
	}

	c.JSON(http.StatusCreated, song)
}
