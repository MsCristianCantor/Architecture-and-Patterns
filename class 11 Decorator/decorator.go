package main

import "fmt"

// Component define la interfaz común para objetos que pueden ser decorados
type Component interface {
	Operation() string
}

// ConcreteComponent es la implementación base de la interfaz Component
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

// Decorator es la clase base para los decoradores, también implementa la interfaz Component
type Decorator struct {
	component Component
}

func (d *Decorator) Operation() string {
	return d.component.Operation()
}

// ConcreteDecoratorA es un decorador concreto que añade funcionalidad extra
type ConcreteDecoratorA struct {
	Decorator
}

func (d *ConcreteDecoratorA) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorA(%s)", d.component.Operation())
}

// ConcreteDecoratorB es otro decorador concreto que añade funcionalidad extra
type ConcreteDecoratorB struct {
	Decorator
}

func (d *ConcreteDecoratorB) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorB(%s)", d.component.Operation())
}

func main() {
	// Crear el componente concreto
	component := &ConcreteComponent{}

	// Decorar el componente con ConcreteDecoratorA
	decoratorA := &ConcreteDecoratorA{
		Decorator{component: component},
	}

	// Decorar el componente con ConcreteDecoratorB sobre el decorador A
	decoratorB := &ConcreteDecoratorB{
		Decorator{component: decoratorA},
	}

	// Ejecutar la operación decorada
	fmt.Println(decoratorB.Operation()) // Salida: ConcreteDecoratorB(ConcreteDecoratorA(ConcreteComponent))
}
