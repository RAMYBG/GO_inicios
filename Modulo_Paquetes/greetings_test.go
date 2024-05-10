package greetings

import (
	"regexp"  // Importa el paquete regexp para trabajar con expresiones regulares
	"testing" // Importa el paquete testing para realizar pruebas unitarias
)

// Test con el nombre
func TestHelloName(t *testing.T) {
	name := "Juan"                                 // Define una variable 'name' con el valor "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`) // Se compila una expresión regular que busca la palabra "Juan" como una palabra completa en un texto. El \b es un metacaracter que representa un límite de palabra, así que esta expresión regular buscará exactamente la palabra "Juan", sin que sea parte de otra palabra más larga.
	msg, err := Hello("Juan")                      // Llama a la función Hello con el argumento "Juan" y captura el mensaje y el error retornados
	if !want.MatchString(msg) || err != nil {      ///// !want.MatchString(msg): Comprueba si el mensaje retornado (msg) no coincide con la expresión regular want. Si la expresión regular no encuentra la palabra "Juan" en msg, la condición se cumple.
		//err != nil: Verifica si err es diferente de nil, lo que indica que hubo un error al ejecutar la función Hello.
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
