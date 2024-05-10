package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especificada

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio") //Regresa un string vacio  un error
	}
	//Hello devuelve un saludo para la persona especificada
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil //Regresa el mensaje y el error
}
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"!Hola,%v! !Bienvenido!",
		"¡Que bueno verte, %v!",
		"¡Saludos, %v! ¡Encantado de conocerte!",
	}
	return formats[rand.Intn(len(formats))]
}
