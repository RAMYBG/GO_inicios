package main

import "fmt"

//Rebanadas En go se pueden agregar y quitar elementos // Arreglos no se puede
func main() {
	var a []int //Hasta aqui declaramos una rebanada
	diasSemana := []string{"Domingo", "Sabado", "Lunes",
		"Martes"} //Rebanada
	diaRebanada := diasSemana[0:2] //De donde a donde de diaSemana//Trozo
	fmt.Println(a, diaRebanada)
	//Para saber cuantos elementos tiene es con len(variable)
	//Para saber cuantos elementos puede contener es con cap si es una rebanada de otra entoces dice
	//de la rebanada tiene 2 pero del total tiene 7 entoces su capacidad es 7
	//make para declarar una rebanada
	nombres := make([]string, 5, 10) //el 5 es el numero de elementos que va almacenar inicialmente y 10 la capacidad
	//Para asignarle valor a la rebana primero es de nombre[0]=Ramiro esto es hasta el indice 4 y despues de eso se debe utilizar append
	nombres[0] = "ramiro"
	fmt.Println(nombres)
	//Fincion copy copia de una rebanada a otra
	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)
	copy(rebanada2, rebanada1)              //Copia de rebanada 1 a la 2
	fmt.Println(copy(rebanada2, rebanada1)) //Imprime la cantidad de elementos que se copiaron
	fmt.Println(rebanada1)
	fmt.Println(rebanada2)
}
