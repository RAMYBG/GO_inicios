package main

import "fmt"

func divide(dividiendo, divisor int) {
	//Funcio anonima
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validateZero(divisor)
	fmt.Println(dividiendo / divisor)
}
func validateZero(divisor int) { //Se puede quitar y el sistema arroja el error
	if divisor == 0 {
		//Se utiliza para un panico
		panic("no se puede dividir por cero")
	}
}

func main() {
	divide(100, 10)
	divide(30, 0)
}
