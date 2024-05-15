package main

import (
	"go-mysql/database"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.COnnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
