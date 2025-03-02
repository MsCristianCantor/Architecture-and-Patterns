package main

import "fmt"

// Observer interface
type Observer interface {
	Update(productName string)
}

// Concrete Observer: Customer
type Customer struct {
	name string
}

func (c *Customer) Update(productName string) {
	fmt.Printf("Hola %s, el producto %s ahora está disponible!\n", c.name, productName)
}

// Subject interface
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify()
}

// Concrete Subject: Product
type Product struct {
	name      string
	available bool
	observers []Observer
}

func NewProduct(name string) *Product {
	return &Product{name: name}
}

func (p *Product) Register(observer Observer) {
	p.observers = append(p.observers, observer)
}

func (p *Product) Unregister(observer Observer) {
	for i, obs := range p.observers {
		if obs == observer {
			p.observers = append(p.observers[:i], p.observers[i+1:]...)
			break
		}
	}
}

func (p *Product) Notify() {
	for _, observer := range p.observers {
		observer.Update(p.name)
	}
}

func (p *Product) SetAvailable() {
	fmt.Printf("\nEl producto %s ahora está disponible!\n", p.name)
	p.available = true
	p.Notify()
}

// Main
func main() {
	// Crear producto
	product := NewProduct("Laptop Gamer")

	// Crear clientes
	customer1 := &Customer{name: "Carlos"}
	customer2 := &Customer{name: "Ana"}

	// Suscribir clientes al producto
	product.Register(customer1)
	product.Register(customer2)

	// El producto se vuelve disponible
	product.SetAvailable()
}
