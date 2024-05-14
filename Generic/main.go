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

func main() {
	//Imprime la lista de valores infinito
	PrintList("Hello", "World", "Go")
	//Imprime la lista de valores infinito
	PrintList(35, "World", 12)
	//Imprime la lista de valores infinito
	PrintList2("Hello", 1, "Go")
	fmt.Printf("Suma: %d \n", Sum(1, 2, 41, 4))
	fmt.Printf("Suma5: %d", Sum5(2, 2, 41, 4))
}
