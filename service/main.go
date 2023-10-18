package main

import (
	"fmt"
	"log"
	"net/http"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprint(w, "Hello World!")
}

func main() {
	handler := http.HandlerFunc(Greeting)
	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
