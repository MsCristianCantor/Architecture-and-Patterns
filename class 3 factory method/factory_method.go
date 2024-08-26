package main

import "fmt"

// Producto es la interfaz que define un comportamiento común
type Product interface {
	Use() string
	Price() string
}

// ConcreteProductA es una implementación concreta de Product
type ConcreteProductA struct{}

func (p ConcreteProductA) Use() string {
	return "Using Product A"
}

func (p ConcreteProductA) Price() string {
	return "Product A price 200"
}

// ConcreteProductB es otra implementación concreta de Product
type ConcreteProductB struct{}

func (p ConcreteProductB) Use() string {
	return "Using Product B"
}

func (p ConcreteProductB) Price() string {
	return "Product B price 300"
}

// ConcreteProductC es otra implementación concreta de Product
type ConcreteProductC struct{}

func (p ConcreteProductC) Use() string {
	return "Using Product C"
}

func (p ConcreteProductC) Price() string {
	return "Product C price 350"
}

// Factory es la interfaz que define el método Factory
type Factory interface {
	CreateProduct() Product
}

// FactoryA es una implementación concreta de Factory que crea ProductA
type FactoryA struct{}

func (f FactoryA) CreateProduct() Product {
	return ConcreteProductA{}
}

// FactoryB es una implementación concreta de Factory que crea ProductB
type FactoryB struct{}

func (f FactoryB) CreateProduct() Product {
	return ConcreteProductB{}
}

// FactoryC es una implementación concreta de Factory que crea ProductC
type FactoryC struct{}

func (f FactoryC) CreateProduct() Product {
	return ConcreteProductC{}
}

func main() {
	var factory Factory

	// Crear un ProductA
	factory = FactoryA{}
	productA := factory.CreateProduct()
	fmt.Println(productA.Use())
	fmt.Println(productA.Price())

	// Crear un ProductB
	factory = FactoryB{}
	productB := factory.CreateProduct()
	fmt.Println(productB.Use())
	fmt.Println(productB.Price())

	// Crear un ProductC
	factory = FactoryC{}
	productC := factory.CreateProduct()
	fmt.Println(productC.Use())
	fmt.Println(productC.Price())
}
