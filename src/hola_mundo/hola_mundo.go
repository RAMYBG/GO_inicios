package main

import (
	"fmt"

	"rsc.io/quote"
)

// Declaracion de constante
const pi float32 = 3.1416

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Hello())
	fmt.Println(pi)
	var firstName, lastName, age = "Alex", "Roe", 27
	println(firstName, lastName, age)
}
