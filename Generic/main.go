package main

import "fmt"

func PrintList(lsit ...interface{}) {
	for _, value := range lsit {
		fmt.Println(value)
	}
}

//Hace lo mismo que interface{}
func PrintList2(lsit ...any) {
	for _, value := range lsit {
		fmt.Println(value)
	}
}

func main() {
	//Imprime la lista de valores infinito
	PrintList("Hello", "World", "Go")
	//Imprime la lista de valores infinito
	PrintList(35, "World", 12)
	//Imprime la lista de valores infinito
	PrintList2("Hello", 1, "Go")
}
