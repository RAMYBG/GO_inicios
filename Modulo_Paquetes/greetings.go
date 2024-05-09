package greetings

import (
	"fmt"
)

// Hello devuelve un saludo para la persona especificada
func Hello(name string) string {
	message := fmt.Sprintf("¡Hola, %v!", name)
	return message
}
