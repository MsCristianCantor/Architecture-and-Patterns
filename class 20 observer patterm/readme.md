# Clase 20: Observer Pattern

## 1. Objetivo
Comprender y aplicar el Observer Pattern, un patrón de diseño de comportamiento que permite definir una dependencia uno a muchos entre objetos, de manera que cuando un objeto cambia de estado, todos sus dependientes son notificados automáticamente.

## 2. Introducción al Observer Pattern
El Observer Pattern es útil cuando se necesita que múltiples objetos reaccionen a cambios en otro objeto sin acoplamiento fuerte. Es común en sistemas de eventos y notificaciones.

### Características clave:
- Define una relación uno a muchos.
- Permite a los observadores suscribirse y desuscribirse dinámicamente.
- Promueve el bajo acoplamiento entre los objetos.

## 3. Estructura del patrón
- **Subject (Sujeto/Observable):** Mantiene una lista de observadores y proporciona métodos para agregar, eliminar y notificar observadores.
- **Observer (Observador):** Define una interfaz para recibir notificaciones.
- **ConcreteSubject:** Implementa la lógica del sujeto y notifica a los observadores cuando cambia de estado.
- **ConcreteObserver:** Implementa la interfaz del observador y actualiza su estado en respuesta a las notificaciones del sujeto.

## 4. Implementación en Go
Implementaremos un ejemplo donde un canal de noticias notifica a sus suscriptores cuando hay una nueva noticia.

```go
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
	name       string
	available  bool
	observers  []Observer
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
```

### Pasos:
1. Definir una interfaz `Observer` con un método `Update()`.
2. Crear la estructura `NewsChannel` como el `Subject`, que maneja una lista de suscriptores.
3. Implementar `Subscribe()` y `Unsubscribe()` en `NewsChannel`.
4. Crear una estructura `Subscriber` que implemente `Observer`.
5. Notificar a los suscriptores cuando haya una nueva noticia.

## 5. Casos de Uso
- Sistemas de eventos y notificaciones.
- Publicación-suscripción en sistemas de mensajería.
- Observadores de estado en interfaces gráficas.

## 6. Ejercicio Práctico
1. Implementa un sistema de notificaciones para una aplicación de clima.
2. Crea un `WeatherStation` (Sujeto) que informe cambios de temperatura.
3. Agrega `TemperatureDisplay` y `MobileApp` como observadores que reaccionen a cambios.
4. Permite suscribirse y desuscribirse dinámicamente.

## 7. Recursos Adicionales
- [Observer Pattern en Go](https://refactoring.guru/design-patterns/observer/go/example)
- [Documentación oficial de Go](https://golang.org/doc/)
- [Ejemplo interactivo en Go Playground](https://go.dev/play/)

