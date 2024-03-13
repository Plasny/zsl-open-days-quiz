package model

import (
	"database/sql"
	_ "embed"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const dbFile = "file:mem_db?mode=memory&cache=shared"

//go:embed schema.sql
var ddl string

var DB *sql.DB

func Init() {
	var err error

	DB, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to SQLite database")

	if _, err := DB.Exec(ddl); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully created database schema")
}
