package main

import (
	"database/sql"

	"log"
)

func connect() *sql.DB {
	//local 127.0.0.1:4003
	//ms 192.168.99.102:4003
	//cloud 35.237.5.35:4003
	db, err := sql.Open("mysql", "root:1234@tcp(192.168.99.102:4003)/project")

	if err != nil {
		log.Fatal("Could not connect to database")
		log.Fatal("No entro base de datos")

	}
	log.Println(db)

	return db
}
