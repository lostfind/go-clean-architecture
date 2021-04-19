package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySqlDB() *sql.DB {
	db, err := sql.Open("mysql", "dwkim:dwkim@tcp(localhost:3306)/zipcode")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
