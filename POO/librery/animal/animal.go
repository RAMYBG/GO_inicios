package animal

import (
	"fmt"
)

type Animal interface {
	Sonido()
}

type Perro struct {
	Nombre string
}
type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + " hace miau")
}
func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + "hace guau guau")
}

func HacerSonido(animal Animal) {
	animal.Sonido()
}
