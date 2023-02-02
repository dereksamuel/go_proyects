package main

import "fmt"

func main() {
	number := 25
	fmt.Println(number)
	fmt.Println(&number) // representacion del valor almacenado en memory
	changeVal(number)
	y := &number
	fmt.Println(y)
	fmt.Println(*y) // ve a esta direccion de memoria y traeme lo que esta almacenado alli
}

func changeVal(a int) { // esta haciendo copia del valor sin hacer el *
	fmt.Println(&a)
	a = 36
}
