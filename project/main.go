package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	// local 8000
	// ms 3003
	log.Println("probando cambios a ver si funcionan ")
	log.Fatal(http.ListenAndServe(":3003", router))
}
