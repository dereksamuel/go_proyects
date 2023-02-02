package main

import "fmt"

type animal interface {
	move() string
}

type lion struct {
	live string
}

type shark struct {
	live string
}

func (lion) move() string {
	return "Walking"
}

func (shark) move() string {
	return "Nadaring"
}

func moveAnimal(a animal) {
	fmt.Println("ANIMAL INTERFACE")
	fmt.Println(a.move())
}

func main() {
	l := lion{}
	s := shark{}

	moveAnimal(l)
	moveAnimal(s)
}
