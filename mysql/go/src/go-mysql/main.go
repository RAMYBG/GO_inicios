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
	handles.ListContacts(db)
}
