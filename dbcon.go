package main

import (
	"database/sql"
	"log"
)

var (
	db  *sql.DB
	err error
)

func dbcon() {
	// Connect to database
	db, err = sql.Open("sqlite3", "./loginsystem.db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	log.Println("db connection successful")
}
