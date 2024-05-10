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

//Aqui se hace la herrrencia
type TextBook struct {
	Book             //Herreda de Book
	editorial string //Agrega mas Valores
	level     string
}

func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}

func (b *TextBook) PrintAll() {
	fmt.Printf("Title: %s\nAuthor: %s\nEditorial:%s\nNivel: %s\n", b.title, b.author, b.pages, b.editorial)
}
