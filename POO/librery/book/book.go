package book

import "fmt"

//Atributos del Libro
//Modelo de un objeto
type Book struct {
	title  string
	author string
	pages  int
}

//Constructor simulacion
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

//Metodo
func (b *Book) PrintInfo() { //Se envia la estructura pero se envia como puntero
	fmt.Printf("Title: %s\nAutor: %s\nPage: %d\n", b.title, b.author, b.pages)

}

//Obtener y Modificar
func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}
