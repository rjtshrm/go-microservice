package service

import (
	"fmt"
	"log"
	"net/http"
)

func returnHelloFromMicroService() string {
	return "Hello From Microservice!"
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request: ", r.Host)
	fmt.Fprint(w, returnHelloFromMicroService())
}

func Server() {
	fmt.Println("Starting Http Server :8080")
	http.HandleFunc("/", responseHandler)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
