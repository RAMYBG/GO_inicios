package main

import (
	"fmt"

	"rsc.io/quote"
)

// Declaracion de constante
const pi float32 = 3.1416
const (
	x = 100
	y = 0b1010 //binario
	z = 0o12   //Octal
	w = 0xFF   //Hexadecimal
)

// Con iota va incrementando de uno en uno y inicia en 0
const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Hello())
	fmt.Println(pi)
	fmt.Println(y, x, w, z)
	var firstName, lastName, age = "Alex", "Roe", 27
	println(firstName, lastName, age)
}
