package main

import "fmt"

// Flyweight Interface
type TreeType interface {
	Display(x, y int)
}

// ConcreteFlyweight
type ConcreteTreeType struct {
	name    string
	color   string
	texture string
}

func (t *ConcreteTreeType) Display(x, y int) {
	fmt.Printf("Árbol %s en color %s mostrado en la posición (%d, %d)\n", t.name, t.color, x, y)
}

// Flyweight Factory
type TreeFactory struct {
	treeTypes map[string]*ConcreteTreeType
}

func NewTreeFactory() *TreeFactory {
	return &TreeFactory{treeTypes: make(map[string]*ConcreteTreeType)}
}

func (f *TreeFactory) GetTreeType(name, color, texture string) *ConcreteTreeType {
	key := name + "-" + color + "-" + texture
	if f.treeTypes[key] == nil {
		f.treeTypes[key] = &ConcreteTreeType{name: name, color: color, texture: texture}
	}
	return f.treeTypes[key]
}

// Cliente
type Tree struct {
	x, y     int
	treeType TreeType
}

func NewTree(x, y int, treeType TreeType) *Tree {
	return &Tree{x: x, y: y, treeType: treeType}
}

func (t *Tree) Display() {
	t.treeType.Display(t.x, t.y)
}

func main() {
	factory := NewTreeFactory()

	// Crear árboles reutilizando Flyweights
	tree1 := NewTree(10, 20, factory.GetTreeType("Pino", "Verde", "Textura1"))
	tree2 := NewTree(30, 40, factory.GetTreeType("Pino", "Verde", "Textura1"))
	tree3 := NewTree(50, 60, factory.GetTreeType("Roble", "Marrón", "Textura2"))

	// Mostrar árboles
	tree1.Display()
	tree2.Display()
	tree3.Display()
}
