#Saludos en go

Este paquete propociona una forma simple de obtener saludos personalizados en go.

##Instalacion 
Ejecuta el siguiente comando para instalar el paquete:
  bash
  go get -u github.com/ramirogarcia/greetings


  ##Uso
  Aqui tines un ejemplo de como utilizar el paquete en tu codigo

  go

  
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
		"Pablo",
		"Juan",
		"Pedro",
	}
	messages, err := greetings.Hellos(names)

	message, err := greetings.Hello("Pablo") //Aqui se manea el error si se envia en vacio genera un error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	fmt.Println(messages)
}
