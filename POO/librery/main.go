package main

import (
	"librery/animal"
)

func main() {
	//Primera forma de crear un objeto sencillo pero mucha linea de codigo
	//Forma dos instanciar y utlizar objeto
	//myBook2 := book.NewBook("Nombre de libro", "Autor del libro", 302)
	//myBook2.PrintInfo()

	//myBook2.SetTitle("Edicion Especial")
	//fmt.Println(myBook2.GetTitle())
	//myBook2.PrintInfo()

	//myTextBook := book.NewTextBook("Comunicacion",
	//	"jaime Gamarra", 263, "Santillana SAC", "Secundaria")
	//myTextBook.PrintInfo()

	//fmt.Print("==========================\n")
	//fmt.Print("==========================\n")
	//fmt.Print("==========================\n")
	//myBook2.PrintInfo()
	//myTextBook.PrintInfo()
	//fmt.Print("==========================\n")
	//fmt.Print("==========================\n")
	//fmt.Print("==========================\n")
	//book.Print(myBook2)
	//book.Print(myTextBook)
	//miPerro := animal.Perro{Nombre: "Max"}
	//miGato := animal.Gato{Nombre: "Tom"}
	//animal.HacerSonido(&miPerro)
	//animal.HacerSonido(&miGato)
	//Se iteran los animales
	animales := []animal.Animal{
		&animal.Perro{
			Nombre: "Max",
		},
		&animal.Gato{
			Nombre: "Tom",
		},
		&animal.Perro{
			Nombre: "Buddy",
		},
		&animal.Gato{
			Nombre: "Milo",
		},
	}
	for _, animal := range animales {
		animal.Sonido()
	}

}
