package main

import "fmt"

// Interfaces para los productos abstractos
type Chair interface {
	SitOn() string
}

type Sofa interface {
	LayOn() string
}

// Productos concretos de la familia 1
type VictorianChair struct{}

func (v VictorianChair) SitOn() string {
	return "Sitting on a Victorian Chair"
}

type VictorianSofa struct{}

func (v VictorianSofa) LayOn() string {
	return "Laying on a Victorian Sofa"
}

// Productos concretos de la familia 2
type ModernChair struct{}

func (m ModernChair) SitOn() string {
	return "Sitting on a Modern Chair"
}

type ModernSofa struct{}

func (m ModernSofa) LayOn() string {
	return "Laying on a Modern Sofa"
}

// Productos concretos de la familia 3
type ArtDecoChair struct{}

func (a ArtDecoChair) SitOn() string {
	return "Sitting on a ArtDeco Chair"
}

type ArtDecoSofa struct{}

func (a ArtDecoSofa) LayOn() string {
	return "Laying on a ArtDeco Sofa"
}

// Abstract Factory
type FurnitureFactory interface {
	CreateChair() Chair
	CreateSofa() Sofa
}

// Fábrica concreta de la familia 1
type VictorianFurnitureFactory struct{}

func (v VictorianFurnitureFactory) CreateChair() Chair {
	return VictorianChair{}
}

func (v VictorianFurnitureFactory) CreateSofa() Sofa {
	return VictorianSofa{}
}

// Fábrica concreta de la familia 2
type ModernFurnitureFactory struct{}

func (m ModernFurnitureFactory) CreateChair() Chair {
	return ModernChair{}
}

func (m ModernFurnitureFactory) CreateSofa() Sofa {
	return ModernSofa{}
}

// Fábrica concreta de la familia 3
type ArtDecoFactory struct{}

func (a ArtDecoFactory) CreateChair() Chair {
	return ArtDecoChair{}
}

func (a ArtDecoFactory) CreateSofa() Sofa {
	return ArtDecoSofa{}
}

func main() {
	// Cliente usando la fábrica Victoriana
	var factory FurnitureFactory
	factory = VictorianFurnitureFactory{}

	chair := factory.CreateChair()
	sofa := factory.CreateSofa()

	fmt.Println(chair.SitOn())
	fmt.Println(sofa.LayOn())

	// Cliente usando la fábrica Moderna
	factory = ModernFurnitureFactory{}

	chair = factory.CreateChair()
	sofa = factory.CreateSofa()

	fmt.Println(chair.SitOn())
	fmt.Println(sofa.LayOn())

	// Cliente usando la fábrica art deco
	factory = ArtDecoFactory{}

	chair = factory.CreateChair()
	sofa = factory.CreateSofa()

	fmt.Println(chair.SitOn())
	fmt.Println(sofa.LayOn())
}
