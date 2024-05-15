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
	// Se crea una instancia de la estructura 'album'.
	var newAlbum album

	// Se intenta vincular (parsear y asignar) los datos JSON recibidos en el cuerpo de la
	// solicitud HTTP a la instancia 'newAlbum'.
	// Si hay un error en este proceso, se termina la función y no se ejecuta el código restante.
	if err := c.BindJSON(&newAlbum); err != nil { //Si no arroja un error se serializa el &newAlbum// De json a type struc albums
		return
	}

	// Se añade el nuevo álbum a la lista global 'albums'.
	albums = append(albums, newAlbum)

	// Se responde al cliente con el código HTTP 201 (Created) y se devuelve la lista actualizada de álbumes en formato JSON.
	c.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return //Si eso pasa termina el codigo
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encotrado"})
}

func main() {
	// Se inicializa el router de Gin.
	router := gin.Default()

	// Se define una ruta GET en '/albums' que ejecuta la función 'getAlbums'.
	router.GET("/albums", getAlbums)

	router.GET("/albums/:id", getAlbumsByID)

	// Se define una ruta POST en '/albums' que ejecuta la función 'postAlbums'.
	router.POST("/albums", postAlbums)

	// Se inicia el servidor en 'localhost' en el puerto 8080.
	router.Run("localhost:8080")
}
