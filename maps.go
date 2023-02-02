package main

import "fmt"

func main() {
	myMap := make(map[string]int) // permite crear el map, map[key]value
	myMap["Derek"] = 89
	fmt.Println(myMap, myMap["Derek"])
}
