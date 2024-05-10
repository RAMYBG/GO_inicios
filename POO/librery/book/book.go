package book // Define el paquete book

import "fmt" // Importa el paquete fmt para la entrada y salida de datos

// Polimorfismo
type Printable interface { // Define una interfaz Printable
	PrintInfo() // Método que deben implementar todos los tipos que sean Printable
}

func Print(p Printable) { // Función que recibe cualquier tipo que implemente la interfaz Printable
	p.PrintInfo() // Llama al método PrintInfo del tipo recibido
}

// Atributos del Libro
// Modelo de un objeto
type Book struct { // Define un tipo struct llamado Book
	title  string // Campo title del tipo string
	author string // Campo author del tipo string
	pages  int    // Campo pages del tipo int
}

// Constructor simulacion
func NewBook(title string, author string, pages int) *Book { // Función que actúa como constructor de un Book
	return &Book{ // Devuelve una nueva instancia de Book
		title:  title,  // Inicializa el campo title
		author: author, // Inicializa el campo author
		pages:  pages,  // Inicializa el campo pages
	}
}

// Metodo
func (b *Book) PrintInfo() { // Método PrintInfo asociado al tipo *Book
	fmt.Printf("Title: %s\nAutor: %s\nPage: %d\n", b.title, b.author, b.pages) // Imprime información del libro
}

// Obtener y Modificar
func (b *Book) SetTitle(title string) { // Método SetTitle para cambiar el título del libro
	b.title = title // Asigna un nuevo título al libro
}

func (b *Book) GetTitle() string { // Método GetTitle para obtener el título del libro
	return b.title // Devuelve el título del libro
}

// Aqui se hace la herencia
type TextBook struct { // Define un tipo struct llamado TextBook
	Book             // Hereda de Book
	editorial string // Agrega un campo adicional llamado editorial
	level     string // Agrega un campo adicional llamado level
}

func NewTextBook(title, author string, pages int, editorial, level string) *TextBook { // Función que actúa como constructor de un TextBook
	return &TextBook{ // Devuelve una nueva instancia de TextBook
		Book:      Book{title, author, pages}, // Inicializa los campos heredados de Book
		editorial: editorial,                  // Inicializa el campo editorial
		level:     level,                      // Inicializa el campo level
	}
}

func (b *TextBook) PrintAll() { // Método PrintAll asociado al tipo *TextBook
	fmt.Printf("Title: %s\nAuthor: %s\nEditorial:%s\nNivel: %d\n", b.title, b.author, b.editorial, b.pages) // Imprime toda la información del libro de texto
}
