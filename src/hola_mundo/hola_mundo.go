package main

import (
	"fmt"

	"rsc.io/quote"
)

// Declaración de constantes
const pi float32 = 3.1416
const (
	x = 100
	y = 0b1010 // Binario
	z = 0o12   // Octal
	w = 0xFF   // Hexadecimal
)

// Uso de iota para enumeración, iniciando en 1
const (
	Domingo = iota + 1
	Lunes
	Martes
	Miércoles
	Jueves
	Viernes
	Sábado
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())
	//Valores predeterminados de cada variable
	var (
		defaultint    int
		defaultUint   uint
		defaultFloat  float32
		defaultBool   bool
		defaultString string
	)
	fmt.Println(defaultint, defaultUint, defaultFloat, defaultBool, defaultString)
}
