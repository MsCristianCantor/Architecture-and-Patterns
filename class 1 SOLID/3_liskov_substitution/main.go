package main

import (
	"fmt"
)

// Bird: Define una interfaz para las aves
type Bird interface {
	Fly()
}

// Sparrow: Implementa la interfaz Bird
type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow flying!")
}

// Penguin: Implementa la interfaz Bird, pero no puede volar
type Penguin struct{}

func (p Penguin) Fly() {
	fmt.Println("Penguins can't fly!")
}

func main() {
	var bird Bird = Sparrow{}
	bird.Fly()

	bird = Penguin{}
	bird.Fly()
}
