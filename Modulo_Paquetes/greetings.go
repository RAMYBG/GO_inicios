package greetings

import (
	"errors"
	"fmt"
)

// Hello devuelve un saludo para la persona especificada
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio") //Regresa un string vacio  un error
	}
	//Hello devuelve un saludo para la persona especificada
	message := fmt.Sprintf("¡ Hola, %v!", name)
	return message, nil //Regresa el mensaje y el error
}
