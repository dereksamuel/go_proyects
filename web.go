package main

import (
	"fmt"
	"io"
	"net/http"
)

type webWriter struct{}

func (webWriter) Write(p []byte) (int, error) {
	// cada uno de los elementos de p es un byte
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")

	if err != nil {
		fmt.Println(err)
	}

	e := webWriter{}
	// fmt.Printf("%T\n", response.Body)

	io.Copy(e, response.Body)
}
