package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// define your data

type album struct {
	ID     string  `json: id`
	Title  string  `json: title`
	Artist string  `json: artist`
	Price  float64 `json: price`
}

// album slide to hold multiple album records

var albumDB = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Main function
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// logics

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albumDB)
}

// add an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON((&newAlbum)); err != nil {
		return
	}

	albumDB = append(albumDB, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// This looks through the DB to get an instance
	for _, a := range albumDB {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(
		http.StatusNotFound,
		gin.H{"message": "album not found"},
	)
}
