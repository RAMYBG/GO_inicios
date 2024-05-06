package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Hello())

	var firstName, lastName, age = "Alex", "Roe", 27
	println(firstName, lastName, age)
}
