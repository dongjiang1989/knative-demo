package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlePost(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "%s", "Hello from Knative Builds!")
}

func main() {
	log.Print("Starting server on port 57890...")
	http.HandleFunc("/", handlePost)
	log.Fatal(http.ListenAndServe(":57890", nil))
}