package main

import "fmt"

type dog struct{}

type fish struct{}

func (dog) walk() (string, string) {
	return "Walking", "I am the dog"
}

func (fish) nadar() (string, string) {
	return "Nadaring", "I am the fish"
}

func moveDog(p dog) {
	fmt.Println(p.walk())
}

func moveFish(p fish) {
	fmt.Println(p.nadar())
}

func main() {
	dog1 := dog{}
	fish1 := fish{}

	moveDog(dog1)
	moveFish(fish1)
}

// Hemos repetido MUCHO CODE
