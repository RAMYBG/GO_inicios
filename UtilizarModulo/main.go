package main

import (
	"fmt"

	"github.com/RAMYBG/GO_inicios"
)

func main() {
	message, err := GO_inicios.Hello("Ramiro")
	if err != nil {
		fmt.Println("Ocurrio un erro: ", err)
	}
	fmt.Println(message)
}
