package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []album{
	{
		ID:     "1",
		Title:  "Familia",
		Artist: "Camila Cabello",
		Year:   2022,
	},
	{
		ID:     "2",
		Title:  "21",
		Artist: "Adele",
		Year:   2011,
	},
	{
		ID:     "3",
		Title:  "The Eminem Show",
		Artist: "Eminem",
		Year:   2022,
	},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func createAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}

func getOneAlbum(value string) *album {
	for _, album := range albums {
		if album.ID == value {
			return &album
		}
	}
	return nil
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	album := getOneAlbum(id)

	if album != nil {
		c.IndentedJSON(http.StatusOK, album)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	oneAlbum := getOneAlbum(id)
	var newAlbum album

	if oneAlbum == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
	}

	for index, a := range albums {
		if oneAlbum.ID == a.ID {
			albums[index] = newAlbum
			if err := c.BindJSON(&newAlbum); err != nil {
				return
			}
		}
	}

	c.IndentedJSON(http.StatusOK, oneAlbum)
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", createAlbum)
	router.GET("albums/:id", getAlbumById)
	router.PUT("albums/:id", updateAlbum)

	if err := router.Run("localhost:8080"); err != nil {
		return
	}
}
