package main

import (
	"fmt"
	"math"

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
	fmt.Println(math.MinInt64, math.MaxInt64)   // Para saber el valor maximo y minimo que puede almacenar
	fullName := "Ramiro \t(alias \"Garcia\")\n" // imprime el nombre parentesis alias "Garcia"parentesis
	fmt.Println(fullName)
	var a byte = 'a'
	fmt.Println(a) //Imprime el valor ascii de a
	s := "hola"
	fmt.Println(s[0]) // Imprime el valor de h en el codigo ascii
}
