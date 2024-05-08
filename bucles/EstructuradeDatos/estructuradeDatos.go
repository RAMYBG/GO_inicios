package main

import "fmt"

func main() {
	var matriz [5]int //Cada valor se consulta por indices
	var a [5]int      //Se le asigno valores a la matriz con ayuda de sus indicess
	a[0] = 20
	a[1] = 11
	var b = [5]int{10, 20, 30, 40} // Tambien se puede declarar asi la matriz
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(matriz)
	var c = [...]int{20, 40, 50} //En este caso es cuando no se sabe el tamaño de la matriz
	fmt.Println(c)
	var matrizbi = [3][3]int{{1, 2, 3}, {3, 2, 1}, {6, 5, 4}}
	fmt.Println(matrizbi)
}
