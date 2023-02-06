package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := 
}

func main() {
	// nombre de la ruta
	http.HandleFunc("/api/v1", index)
	fmt.Println("Hello world")

	// start with the server
	http.ListenAndServe(":3500", nil)
}
