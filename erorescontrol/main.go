package main

import (
	"errors" //Utilizar biblioteca errors
	"fmt"
	"strconv"
)

func divide(dividiendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por cero") //Declaracio de nuestros propios errores
	}
	return dividiendo / divisor, nil
}

func main() {
	//Error al dividir
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Resultado:", result)
	//Error al convertir un string a un entero
	str := "123f"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return //Termina la funcion
	}
	fmt.Println("Numero:", num)
}
