package services

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:P@sword@tcp(127.0.0.1:3306)/livecode")
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db, err
}
