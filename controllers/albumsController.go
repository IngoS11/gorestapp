package controllers

import (
	"net/http"
	"strconv"

	"github.com/IngoS11/gorestapp/initializers"
	"github.com/IngoS11/gorestapp/models"
	"github.com/gin-gonic/gin"
)

type Album struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

// AddAlbum godoc
//
//	@Summary		add an album
//	@Description	add an album by posting json
//	@Tags			albums
//	@Accept			json
//	@Produc			json
//	@Param			album	body		controllers.Album	true	"Add model"
//	@Success		200		{object}	controllers.Album
//	@Router			/albums [post]
func AddAlbum(c *gin.Context) {
	// get the title, artist, price off the request body
	var body Album

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	album := models.Album{Title: body.Title, Artist: body.Artist, Price: body.Price}
	// create the album
	result := initializers.DB.Create(&album)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to add Album to database",
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}

// GetAlbumById godoc
//
//	@Summary		Get an album by it's id
//	@Description	Returns an album by it's id
//	@Tags			albums
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int		true	"Album ID"
//	@Success		200	{string}	string	"ok"
//	@Router			/albums/{id} [get]
func GetAlbumById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var body Album
	var album models.Album
	initializers.DB.First(&album, id)
	if album.ID == 0 {
		c.IndentedJSON(http.StatusNotFound, id)
		return
	}
	body.Artist = album.Artist
	body.Price = album.Price
	body.Title = album.Title
	c.IndentedJSON(http.StatusOK, body)
}

// GetAllAlbums godoc
//
//	@Summary		get all albums
//	@Description	returns all albums in the system
//	@Tags			albums
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"ok"
//	@Router			/albums [get]
func GetAllAlbums(c *gin.Context) {

	var albums []models.Album
	var line Album
	var body []Album
	result := initializers.DB.Find(&albums)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retreive albums from database",
		})
	}

	for _, album := range albums {
		line.Artist = album.Artist
		line.Price = album.Price
		line.Title = album.Title
		body = append(body, line)
	}
	c.JSON(http.StatusOK, body)
}
