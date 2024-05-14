package main

import (
	"github.com/gin-gonic/gin"
)

type album struct {
	ID      string  `json:"id`
	Title   string  `json:"title"`
	Artista string  `json:"artista"`
	Price   float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artista: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artista: "gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artista: "Sarah Vaughan", Price: 15.99},
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.Run("localhost:8080")
}
