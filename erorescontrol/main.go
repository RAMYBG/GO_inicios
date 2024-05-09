package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Estructura de contactos
type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

// Guarda contactos en un archivo json
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file) //Es para convertir estructura de datos en formato json
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

// Carga contactos desde un archivo json
func loadContasctsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decode := json.NewDecoder(file)
	err = decode.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	//Slice de contactos
	var contacts []Contact
	//Cargar contactos existentes desde el archivo
	err := loadContasctsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al carga los contactos", err)
	}
	//Crear instancia de fubio
	reader := bufio.NewReader(os.Stdin)

	for {
		//Mostrar opciones al usuario
		fmt.Print("==== GESTOR DE CONTACTOS ====\n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Elige una opcion: ")
		//Leer la opcion del usuario
		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion", err)
			return
		}
		switch option {
		case 1:
			//Ingresar y crear contactos
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			//Agregar un contacto a Slice
			contacts = append(contacts, c)
			//Guardar en un archivo json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar los contactos", err)
			}
		case 2:
			//Mostrar todos los contactos
			fmt.Println("========================================")
			for index, contact := range contacts {
				fmt.Printf("%d . Nombre: %s Email: %s Telefono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("========================================")
		case 3:
			//Salir del programa
			return
		default:
			fmt.Println("Opcion no valida")
		}

	}
}
