package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func (s *Persona) diHola() { //Se utiliza el puntero de la estructura Persona s es la instancia probicional
	fmt.Println("Hola, mi nombre es", s.nombre)
}

//Con el puntero se hace referencia a la memoria
func main() {
	s := Persona{"Ramiro", 23, "correo@gmail.com"}
	s.diHola()
	var x int = 10
	var p *int = &x
	fmt.Println(p)
	fmt.Println(&x)
	fmt.Println(*p)
	editar(&x)
	fmt.Println(x)
}

func editar(x *int) {
	*x = 20
}
