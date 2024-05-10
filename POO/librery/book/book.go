package book

import "fmt"

//Atributos del Libro
//Modelo de un objeto
type Book struct {
	Title  string
	Author string
	Pages  int
}

//Constructor simulacion
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

//Metodo
func (b *Book) PrintInfo() { //Se envia la estructura pero se envia como puntero
	fmt.Printf("Title: %s\nAutor: %s\nPage: %d\n", b.Title, b.Author, b.Pages)

}
