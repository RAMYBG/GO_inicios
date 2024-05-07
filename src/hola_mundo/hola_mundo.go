package main

import (
	"fmt"
	"math"

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
	fmt.Println(quote.Hello())
	fmt.Println(pi)
	fmt.Println(y, x, w, z)
	var firstName, lastName, age = "Alex", "Roe", 27
	println(firstName, lastName, age)
	fmt.Println(math.MinInt64, math.MaxInt64)   // Valores máximo y mínimo para int64
	fullName := "Ramiro \t(alias \"Garcia\")\n" // Nombre y alias en formato tabulado
	fmt.Println(fullName)

	//TIpo de dato byte en ascii
	var a byte = 'a'
	fmt.Println(a) // Imprime el valor ASCII de 'a'
	s := "hola"
	fmt.Println(s[0]) // Imprime el valor ASCII de 'h'
	// Uso de tipo de dato rune para almacenar un caracter Unicode
	var r rune = '❤' // Imprime el unicode o ascii del corazon
	fmt.Println(r)
}
