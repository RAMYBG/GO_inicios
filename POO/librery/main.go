package main

import "librery/book"

func main() {
	//Primera forma de crear un objeto sencillo pero mucha linea de codigo
	var myBook = book.Book{ //Se instancia
		Title:  "The Go Programming Language",
		Author: "John W. Johnson",
		Pages:  23,
	}
	myBook.PrintInfo() //Se utiliza el metodo
	//Forma dos instanciar y utlizar objeto
	myBook2 := book.NewBook("Nombre de libro", "Autor del libro", 302)
	myBook2.PrintInfo()

}
