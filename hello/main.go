package main

import (
	"fmt"
	"log"

	"github.com/ramirogarcia/greetings"
)

func main() {
	log.SetPrefix("greetings") //prefijo en los mensaje de registro
	log.SetFlags(0)            //Bandera de formato en 0// No se mostrar la fecha y hora

	names := []string{
		"Ramiro",
		"Juan",
		"Pedro",
	}
	messages, err := greetings.Hellos(names)

	message, err := greetings.Hello("Ramiro") //Aqui se manea el error si se envia en vacio genera un error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	fmt.Println(messages)
}
