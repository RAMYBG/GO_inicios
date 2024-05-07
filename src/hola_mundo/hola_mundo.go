package main

import (
	"fmt"
	"strconv"

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
	var interger16 int16 = 50
	var integer32 int32 = 100
	//Para sumar valores hay que comvertirlos en el mismo formato int32 o int16
	fmt.Println(interger16 + int16(integer32))

	//Cuando se quiere convertir un string a un entero se necesita de un paquete
	s := "100"
	//i, _ := strconv.Atoi(s) //Convierte el string en entero _ espara no utilizar el error

	n := 42
	s = strconv.Itoa(n)
	fmt.Println(s + s)
	name := "Ramiro"
	age := 23
	//TIpos de Print con formato
	fmt.Printf("Hola, me llamo %s y tengo %d años.\n", name, age)
	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.\n", name, age)
	fmt.Println(greeting)
	///////
	/////////
	var name1 string
	var name2 string
	var age1 int
	//Para insertar datos desde la co sola
	fmt.Println("Ingrese su nombre:")
	fmt.Scanln(&name1, &name2)
	fmt.Println("Ingrese su edad:")
	fmt.Scanln(&age1)
	fmt.Printf("Hola, me llamo %s %sy tengo %d años.\n", name1, name2, age1)
	//Para regresar el tipo de dato
	fmt.Printf("El tipo de name es: %T\n", name1)
}
