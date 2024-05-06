package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Hello())

	var (
		firstName, lastName string
		age                 int
	)
	firstName = "Alex"
	lastName = "Roe"
	age = 27
	println(firstName, lastName, age)
}
