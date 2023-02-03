package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	started := time.Now()

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, server := range servers {
		checkServer(server)
	}

	finished := time.Since(started)

	fmt.Printf("Current time is: %s\n", finished)
}

func checkServer(server string) {
	_, err := http.Get(server)

	if err != nil {
		fmt.Println("[server-error]:", err)
		return
	}

	fmt.Printf("[server-%s-ok]\n", server)
}

// LO ESTA HACIENDO DE FORMA SECUENCIAL Y NUESTRO PROGRAMA QUEDA BLOQUEADO y desperdiciamos valioso time, go run servers.go
