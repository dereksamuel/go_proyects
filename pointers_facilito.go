package main

import "fmt"

func main() {
	/*
		1. a pointer es una address in memory
		2. *T => *int, *string, *Struct
	*/

	var x, y *int
	integer := 5

	x = &integer // un valor en memoria de direccion en memoria

	fmt.Println(*x, y)
}
