package main

import "fmt"

//Funcion variadica
func suma(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	suma(12, 3, 4, 364, 2, 12)
	fmt.Println(suma(102, 34, 52, 44))
}
