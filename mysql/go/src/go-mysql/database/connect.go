package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func COnnect() (*sql.DB, error) {
	//Parametro para conexion a la base de datos
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	//Abrir una conexion a la base de datos
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Conexion de la base de datos exitosa")
	return db, nil
}
