package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

// Estructura de datos
func main() {
	var persona1 Persona //Se instancia
	persona1.nombre = "Ramiro"
	persona1.edad = 23
	persona1.correo = "ramiro@gamil.com"
	fmt.Println(persona1)
	//Otra forma de instanciar
	persona2 := Persona{"ramiro", 23, "Correo@gmail.com"}
	fmt.Println(persona2)
}
