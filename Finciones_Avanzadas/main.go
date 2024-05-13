package main

import "fmt"

//Funcion variadica
func suma(name string, nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Hola %s, la suma es %d\n", name, total)
	return total
}

func imprimirDatos(datos ...interface{}) { //Interface es para recibir dats sin discriminacion
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

func main() {
	fmt.Println(suma("Ramiro", 102, 34, 52, 44))
	imprimirDatos("Hola", 132, 123.45, true) //Imprime que tipo de datos es y recibe N cantidad de datos sin importar el tipo
}
