package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	leer := bufio.NewReader(os.Stdin) //Esto es para leer la entrada de varios caracteres
	//en el scanln
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
		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingresa el nombre de la tarea:")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente.")
		case 2:
			var index int
			fmt.Print("Ingrese el indice de la tarea a marcar como completada: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada")
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el indice de la tarea a actuaiizar: ")
			fmt.Scanln(&index)
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea actualizada correctamente.")
		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea a eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada correctamente.")
		case 5:
			fmt.Print("Saliendo del programa...")
			return
		default:
			fmt.Println("Opcion no valida")
			//Listar todas las tareas
			fmt.Println("Lista de tareas")
			fmt.Println("===============================")
			for i, t := range lista.tareas {
				fmt.Printf("%d. %s -%s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
			}
			fmt.Println("================================")
		}
	}

}
