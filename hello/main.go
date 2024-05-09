package main

import (
	"fmt"

	"github.com/ramirogarcia/greetings"
)

func main() {
	message := greetings.Hello("Ramiro")
	fmt.Println(message)
}
