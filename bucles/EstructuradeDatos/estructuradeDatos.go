package main

import "fmt"

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

// Metodo para agregar tarea
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// Metodo para marcar como completado
func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

// metodo para editar tarea
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

// Metodo para eliminar tarea
func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}
func main() {
	//Instancia de lista d etareas
	lista := ListaTareas{}
	for {
		var opcion int
		fmt.Println("Selecione una opcion:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir")
		fmt.Println("Ingrese la opcion: ")
		fmt.Scanln(&opcion)
	}
	//Listar todas las tareas
	fmt.Println("Lista de tareas")
	fmt.Println("===============================")
	for i, t := range lista.tareas {
		fmt.Printf("%d. %s -%s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
	}
	fmt.Println("================================")
}
