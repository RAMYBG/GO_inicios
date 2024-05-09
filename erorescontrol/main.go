package main

import (
	"fmt"
	"os"
)

func main() {
	//Se ejecuta de orden jerarquico por el defer
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println(1)
	//Se ejecuta de orden jerarquico por el defer

	file, err := os.Create("hola.txt") //Es para crear el archivo y ontener su errror si ocurre
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() //Se ejecuta antes de que se termine de ejecutar el main
	_, err = file.Write([]byte("Hola, Ramiro Garcia"))
	if err != nil {
		fmt.Println(err)
		return
	}

}
