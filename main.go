package main

import (
	"fmt"
	"log"
	"net/http"
)

func responseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello From Microservice!")
}

func main() {
	fmt.Println("Starting Http Server :8080")
	http.HandleFunc("/", responseHandler)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
