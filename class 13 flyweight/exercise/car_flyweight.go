package main

import "fmt"

// Flyweight Interface
type CarType interface {
	StreetPosition(x, y int)
}

// ConcreteFlyweight
type ConcreteCarType struct {
	carType string
	color   string
	mark    string
}

func (t *ConcreteCarType) StreetPosition(x, y int) {
	fmt.Printf(
		"Carro de marca %s tipo %s en color %s mostrado en la posicion (%d, %d) de la calle\n",
		t.mark,
		t.carType,
		t.color,
		x,
		y,
	)
}

// Flyweight Factory
type CarFactory struct {
	carTypes map[string]*ConcreteCarType
}

func NewCarFactory() *CarFactory {
	return &CarFactory{carTypes: make(map[string]*ConcreteCarType)}
}

func (f *CarFactory) GetCarType(carType, color, mark string) *ConcreteCarType {
	key := carType + "-" + color + "-" + mark
	if f.carTypes[key] == nil {
		f.carTypes[key] = &ConcreteCarType{carType: carType, color: color, mark: mark}
	}
	return f.carTypes[key]
}

// Cliente
type Car struct {
	x, y    int
	carType CarType
}

func NewCar(x, y int, carType CarType) *Car {
	return &Car{x: x, y: y, carType: carType}
}

func (t *Car) StreetPosition() {
	t.carType.StreetPosition(t.x, t.y)
}

func main() {
	factory := NewCarFactory()

	// Crear autos reutilizando Flyweights
	car1 := NewCar(10, 20, factory.GetCarType("sedan", "Verde", "renault"))
	car2 := NewCar(30, 40, factory.GetCarType("QP", "rojo", "BMW"))
	car3 := NewCar(50, 60, factory.GetCarType("sedan", "Verde", "renault"))

	// Mostrar autos
	car1.StreetPosition()
	car2.StreetPosition()
	car3.StreetPosition()
}
