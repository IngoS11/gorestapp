package api

import (
	"net/http"
	"strconv"

	"github.com/IngoS11/gorestapp/model"
	"github.com/gin-gonic/gin"
)

// GetAlbum godoc
//
//	@Summary		get all albums
//	@Description	returns all albums in the system
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"ok"
//	@Router			/albums [get]
func GetAlbum(c *gin.Context) {
	albums, err := model.AllAlbums()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, albums)
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// GetAlbumById godoc
//
//	@Summary		Get an album by it's id
//	@Description	Returns an album by it's id
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int		true	"Album ID"
//	@Success		200	{string}	string	"ok"
//	@Router			/albums/{id} [get]
func GetAlbumById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	album, err := model.AlbumsByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, id)
	}
	c.IndentedJSON(http.StatusOK, album)
}

// PostAlbum godoc
//
//	@Summary		add an album
//	@Description	add an album by posting json
//	@Tags			albums
//	@Accept			json
//	@Produc			json
//	@Param			album	body		model.Album	true	"Add model"
//	@Success		200		{object}	model.Album
//	@Router			/albums [post]
func PostAlbum(c *gin.Context) {
	var newAlbum model.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the database
	albumId, err := model.AddAlbum(newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, albumId)
	}
	c.IndentedJSON(http.StatusOK, albumId)

}
