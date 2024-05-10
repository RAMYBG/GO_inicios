package main

import (
	"fmt"
	"librery/book"
)

func main() {
	//Primera forma de crear un objeto sencillo pero mucha linea de codigo
	//Forma dos instanciar y utlizar objeto
	myBook2 := book.NewBook("Nombre de libro", "Autor del libro", 302)
	myBook2.PrintInfo()

	myBook2.SetTitle("Edicion Especial")
	fmt.Println(myBook2.GetTitle())

}
