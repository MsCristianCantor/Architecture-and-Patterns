package main

import "fmt"

// Producto final que queremos construir
type House struct {
	Foundation string
	Structure  string
	Roof       string
	Interior   string
	Color      string
}

// Builder es la interfaz que define los pasos para construir la casa
type Builder interface {
	SetFoundation()
	SetStructure()
	SetRoof()
	SetInterior()
	SetColor()
	GetHouse() House
}

// ConcreteBuilder es una implementación concreta del Builder
type ConcreteBuilder struct {
	house House
}

func (b *ConcreteBuilder) SetFoundation() {
	b.house.Foundation = "Concrete, brick, and stone"
}

func (b *ConcreteBuilder) SetStructure() {
	b.house.Structure = "Wood and bricks"
}

func (b *ConcreteBuilder) SetRoof() {
	b.house.Roof = "Concrete and steel"
}

func (b *ConcreteBuilder) SetInterior() {
	b.house.Interior = "Wood and paint"
}

func (b *ConcreteBuilder) SetColor() {
	b.house.Color = "Color White"
}

func (b *ConcreteBuilder) GetHouse() House {
	return b.house
}

// ConcreteBuilder es una implementación concreta del Builder
type ConcreteModernBuilder struct {
	house House
}

func (b *ConcreteModernBuilder) SetFoundation() {
	b.house.Foundation = "Concrete, brick, and marble"
}

func (b *ConcreteModernBuilder) SetStructure() {
	b.house.Structure = "Wood and bricks"
}

func (b *ConcreteModernBuilder) SetRoof() {
	b.house.Roof = "Concrete and steel"
}

func (b *ConcreteModernBuilder) SetInterior() {
	b.house.Interior = "Marble"
}

func (b *ConcreteModernBuilder) SetColor() {
	b.house.Color = "Color Black"
}

func (b *ConcreteModernBuilder) GetHouse() House {
	return b.house
}

// Director es responsable de manejar el proceso de construcción
type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(b Builder) {
	d.builder = b
}

func (d *Director) ConstructHouse() {
	d.builder.SetFoundation()
	d.builder.SetStructure()
	d.builder.SetRoof()
	d.builder.SetInterior()
	d.builder.SetColor()
}

func main() {
	// Crear el Builder concreto y el Director
	builder := &ConcreteBuilder{}
	director := Director{builder: builder}

	// Construir la casa
	director.ConstructHouse()

	// Obtener el producto final
	house := builder.GetHouse()

	fmt.Printf("House built with: %s, %s, %s, %s, %s\n",
		house.Foundation, house.Structure, house.Roof, house.Interior, house.Color)

	modernBuilder := &ConcreteModernBuilder{}
	director.SetBuilder(modernBuilder)

	// Construir la casa moderna
	director.ConstructHouse()

	// Obtener el producto final
	modernHouse := modernBuilder.GetHouse()

	fmt.Printf("Modern House built with: %s, %s, %s, %s, %s\n",
		modernHouse.Foundation, modernHouse.Structure, modernHouse.Roof, modernHouse.Interior, modernHouse.Color)
}
