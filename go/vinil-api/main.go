package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Se creo la estructura de los datos
type album struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Artista string  `json:"artista"`
	Price   float64 `json:"price"`
}

// Se declara una lista de datos
var albums = []album{
	{ID: "1", Title: "Blue Train", Artista: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artista: "gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artista: "Sarah Vaughan", Price: 15.99},
}

// Creamos un controlador
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) //Esto hace respoder al cliente con la data a tipo json y el estatus http

}

func postAlbums(c *gin.Context) {
	//se crea una instancia
	var newAlbum album
	// Se lee el body del request
	//Se delcra una condicion
	//Se vinxula los datos que recibe al album
	if err := c.BindJSON(&newAlbum); err != nil { //Recibe una estructura y devuelve un error
		return
	}
	albums = append(albums, newAlbum)
	//No cambia al anterior
	c.IndentedJSON(http.StatusCreated, albums)
}
func main() {
	router := gin.Default()
	//Metodo GET Trae todos los elementos de la lista
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	//Inicia el servidsor
	router.Run("localhost:8080")
}
