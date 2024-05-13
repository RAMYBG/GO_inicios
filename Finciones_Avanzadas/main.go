// Declara el nombre del paquete en el que está definido este código. En Go, todo código pertenece a un paquete.
package main

// Importa el paquete "fmt", que se utiliza para operaciones de entrada y salida, como imprimir texto en la consola.
import "fmt"

// Define una función llamada 'factorial' que toma un entero 'n' como argumento y retorna un entero.
func factorial(n int) int {
	// Condición base de la recursión: si n es 0, retorna 1, porque el factorial de 0 es 1.
	if n == 0 {
		return 1
	}
	// Caso recursivo: retorna n multiplicado por el factorial de n-1.
	return n * factorial(n-1)
}

// La función 'main' es el punto de entrada del programa. En Go, la ejecución comienza con main.
func main() {
	// Imprime el resultado de la función 'factorial' con el argumento 10, usando la función Println del paquete 'fmt'.
	fmt.Println(factorial(10))
}
