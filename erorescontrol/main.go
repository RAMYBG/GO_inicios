package main

import "log"

func main() {
	log.SetPrefix("main: ")                         //Prefijo para las lineas siguientes
	log.Print("Este es un mesaje de registro")      //Inprime el tiempo en que ocurio
	log.Println("Este es otro mensaje de registro") //Hace lo mismo que el anterior basicamente
	//log.Fatal("Este es otro mensaje de registro")   // Se ejecuta y detiene la ejecucion  el programa
	log.Panic("Este es otro mensaje de registro") //Nos muestra una pila de errores y registro y tambien termina la ejecucion

}
