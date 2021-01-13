package main

import (
	"log"
	"net/http"
)

func main() {
	router := RouterApp()

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}


func ShowError(e error) {
	if e != nil { // nil = without error
		log.Fatal(e)
	}
}