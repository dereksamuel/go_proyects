package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	started := time.Now()
	channel := make(chan string)

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	var i int

	for {
		if i >= len(servers)*2 {
			break
		}

		for _, server := range servers {
			go checkServerGo(server, channel)
		}

		time.Sleep(1 * time.Second)
		fmt.Println(<-channel)
		i++
	}

	// for i := 0; i < len(servers); i++ {
	// 	fmt.Println(<-channel)
	// } // va a esperar a que todos los goroutines terminen
	finished := time.Since(started)

	fmt.Printf("Current time is: %s\n", finished)
}

func checkServerGo(server string, channel chan string) {
	_, err := http.Get(server)

	if err != nil {
		channel <- "[server-" + server + "-ERROR]"
		return
	}

	channel <- "[server-" + server + "-OK]"
}
