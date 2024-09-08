package main

import "fmt"

type Component interface {
	Operation() string
	Add(Component)
	Remove(Component)
	GetChild(int) Component
}

type Leaf struct {
	name string
}

func (l *Leaf) Operation() string {
	return l.name
}

func (l *Leaf) Add(c Component) {
	// No hace nada ya que es una hoja
}

func (l *Leaf) Remove(c Component) {
	// No hace nada ya que es una hoja
}

func (l *Leaf) GetChild(i int) Component {
	return nil // Las hojas no tienen hijos
}

type Composite struct {
	children []Component
}

func (c *Composite) Operation() string {
	result := "Composite ["
	for _, child := range c.children {
		result += child.Operation() + " "
	}
	result += "]"
	return result
}

func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Remove(component Component) {
	// Implementar lógica para eliminar el componente
}

func (c *Composite) GetChild(i int) Component {
	if i >= 0 && i < len(c.children) {
		return c.children[i]
	}
	return nil
}

func main() {
	leaf1 := &Leaf{name: "Leaf 1"}
	leaf2 := &Leaf{name: "Leaf 2"}
	composite := &Composite{}
	composite.Add(leaf1)
	composite.Add(leaf2)

	fmt.Println(composite.Operation()) // Salida: Composite [Leaf 1 Leaf 2 ]
}
