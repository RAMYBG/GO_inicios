package main

import (
	"fmt"
	//Se utiliza el modulo de paquetes de GO
	"golang.org/x/exp/constraints"
)

func PrintList(lsit ...interface{}) {
	for _, value := range lsit {
		fmt.Println(value)
	}
}

// Hace lo mismo que interface{}
func PrintList2(lsit ...any) {
	for _, value := range lsit {
		fmt.Println(value)
	}
}

// Aproximacion
func Sum[T ~int | ~float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Restrinccion union de elementor
func Sum2[T int | float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Restrinccion albritrario
// Restrinccion
func Sum3[T int](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

type Numers interface {
	~int | ~float64 | ~float32 | ~uint
}

// Restrinccion albritrario
// Restrinccion
func Sum4[T Numers](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Se utiliza el  modulo
func Sum5[T constraints.Integer](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// Crea una funcion para filtrar utiliza la restrinccion de Oredered recibe una lista con la restriccion, recibe una funcion que recibe un entero o los dstos que este marcado la restriccion y regresa un bool
func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	resul := make([]T, 0, len(list)) //Crea una lista para guardar el filtro despues
	for _, item := range list {
		if callback(item) { //Si la funcion callback devuelve verdadero
			resul = append(resul, item) //Se hace apppend a la lista resul
		}
	}
	return resul
}

type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}

func main() {
	lista1 := []string{"a", "b", "c"}
	lista2 := []int{1, 2, 3, 4, 5}
	//Imprime la lista de valores infinito
	PrintList("Hello", "World", "Go")
	//Imprime la lista de valores infinito
	PrintList(35, "World", 12)
	//Imprime la lista de valores infinito
	PrintList2("Hello", 1, "Go")
	fmt.Printf("Suma: %d \n", Sum(1, 2, 41, 4))
	fmt.Printf("Suma5: %d\n", Sum5(2, 2, 41, 4))
	fmt.Println(Includes(lista1, "g"))
	fmt.Println(Includes(lista1, "a"))
	fmt.Println(Includes(lista2, 1))
	fmt.Println(Includes(lista2, 8))

	//Comienza a crear la lista apartir de la condicion del return
	fmt.Println(Filter(lista2, func(value int) bool { return value > 3 }))

	/////////////////////
	////////////////
	//Crear un producto //estructura Generica
	product1 := Product[uint]{1, "Zapato", 50}
	product2 := Product[string]{"SNDVNDFVDVSFDGDF-DG-DFH-DSV", "Zapato", 50}
	fmt.Println(product1)
	fmt.Println(product2)
}
