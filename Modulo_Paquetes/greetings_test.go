package greetings

import (
	"regexp"  // Importa el paquete regexp para trabajar con expresiones regulares
	"testing" // Importa el paquete testing para realizar pruebas unitarias
)

// Test con el nombre
func TestHelloName(t *testing.T) {
	name := "Juan"                                 // Define una variable 'name' con el valor "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`) // Compila una expresión regular que busca la palabra "Juan" como palabra completa
	msg, err := Hello("Juan")                      // Llama a la función Hello con el argumento "Juan" y captura el mensaje y el error retornados
	if !want.MatchString(msg) || err != nil {      // Verifica si el mensaje no coincide con la expresión regular o si hay un error
		t.Fatalf(`Hello("Juan")= %q, %v, quiere coincidencia para %#q, nil`, msg, err, want) // Registra un error fatal en la prueba si no se cumplen las condiciones esperadas
	}
}

// Test vacío
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")        // Llama a la función Hello con un argumento vacío y captura el mensaje y el error retornados
	if msg != "" || err == nil { // Verifica si el mensaje no está vacío o si no hay un error
		t.Fatalf(`Hello("")= %q, %v, quiere "", error`, msg, err) // Registra un error fatal en la prueba si no se cumplen las condiciones esperadas
	}
}
