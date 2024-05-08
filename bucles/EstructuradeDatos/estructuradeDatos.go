package main

import "fmt"

//Rebanadas En go se pueden agregar y quitar elementos // Arreglos no se puede
func main() {
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00F00",
		"azul":  "#0000FF",
	}
	fmt.Println(colors["rojo"]) //Imprime el valor hexadecimal
	colors["negro"] = "000000"
	fmt.Println(colors)
	if valor, ok := colors["blanco"]; ok { //Se decalra la variable valor y tambien se hace la autentificacion y despues se evalua el ok =verdadero en el if
		fmt.Println(valor)
	} else {
		fmt.Println("No existe la clave")
	}
	delete(colors, "azul") //Elimina el azul dentro del mapa
	fmt.Println(colors)

	//Iterar todos los valores del map
	for clave, valor := range colors { //Clave nombre, valor valor ...map[string]string
		fmt.Printf("Clave: %s, valor: %s\n", clave, valor) //Formato
	}
}
