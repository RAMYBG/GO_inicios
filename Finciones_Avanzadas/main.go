package main

import "fmt"

func saludar(name string, f func(string)) {
	f(name)
}

func main() {
	//Declaracion de una funcio anonima
	//func() {
	//	fmt.Println("Hola soy una funcion anonima")
	//}()
	//Otra opcion en donde se puede declarar una varible a una funcion anonima
	saludo := func(name string) {
		fmt.Printf("Hola soy %s", name)
	}
	//saludo("ramiro")
	//Otra opcion en donde se puede declarar una varible a una funcion anonima
	saludar("Ramiro", saludo) //La funcion normal saludar recibe el string y tambie recibe una funcion anonima
	//la cual esta dentro de una variable
}
