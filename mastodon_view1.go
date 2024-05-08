package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":8081"
	log.Printf("Starting view1 at port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
