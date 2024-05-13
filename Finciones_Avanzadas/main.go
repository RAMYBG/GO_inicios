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

func main() {
	fmt.Println(suma("Ramiro", 102, 34, 52, 44))
}
