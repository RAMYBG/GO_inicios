package main

import (
	"go-mysql/database"
	"go-mysql/handles"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.COnnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//handles.GetContactByID(db, 3)

	//Prepara los datos que se van a registrar en l abse de datos con el modelo
	//newContact := models.Contact{
	//	Name:  "Nuevo Usuario",
	//Email: "nuevo@example.com",
	//	Phone: "123463",
	//}
	//Datos para actualizar con el ID
	//newContact2 := models.Contact{
	//Id:    1,
	//Name:  "Usuario",
	//Email: "nuevo@example.com",
	//Phone: "1289575863",
	//}
	//Utiliza la funcion para hacer la insercion
	//handles.CreateContac(db, newContact)
	//handles.UpdateCOntact(db, newContact2)
	//Elimina un registro
	handles.DeleteContact(db, 5)
	//Lista todos los contactos
	handles.ListContacts(db)
}
