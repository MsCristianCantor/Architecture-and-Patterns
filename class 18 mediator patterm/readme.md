# Clase 18: Mediator Pattern

## 1. Objetivo
Comprender el **patrón Mediator** y su aplicación en sistemas que requieren comunicación descentralizada entre objetos. Implementar este patrón en Go para facilitar la comunicación entre componentes sin generar dependencias fuertes.

---

## 2. Introducción
El **patrón Mediator** es un patrón de diseño de comportamiento que encapsula la comunicación entre objetos en un componente central denominado *mediador*.

Este patrón se utiliza para evitar la comunicación directa entre objetos y reducir la dependencia mutua entre ellos, lo que facilita el mantenimiento y la extensibilidad del sistema.

**Ejemplo típico:** Un chat donde los usuarios envían mensajes a través de un servidor central en lugar de comunicarse directamente entre ellos.

---

## 3. Estructura del Patrón
El patrón Mediator está compuesto por:

- **Mediador (Mediator)**: Define una interfaz para la comunicación entre los objetos.
- **Mediador Concreto (Concrete Mediator)**: Implementa la interfaz del mediador y coordina la comunicación entre los objetos.
- **Colega (Colleague)**: Representa los objetos que interactúan con el mediador en lugar de comunicarse entre sí directamente.

**Diagrama UML:**

![mediator.jpg](mediator.jpg)

---

## 4. Implementación en Go
A continuación, se presenta una implementación básica del patrón Mediator en Go:

```go
package main

import "fmt"

// Mediator interface
type ChatMediator interface {
	SendMessage(msg string, user User)
	AddUser(user User)
}

// Concrete Mediator
type ChatRoom struct {
	users []User
}

func (c *ChatRoom) SendMessage(msg string, user User) {
	for _, u := range c.users {
		if u != user {
			u.ReceiveMessage(msg)
		}
	}
}

func (c *ChatRoom) AddUser(user User) {
	c.users = append(c.users, user)
}

// Colleague interface
type User interface {
	SendMessage(msg string)
	ReceiveMessage(msg string)
}

// Concrete Colleague
type ChatUser struct {
	name string
	mediator ChatMediator
}

func (u *ChatUser) SendMessage(msg string) {
	fmt.Printf("%s envía: %s\n", u.name, msg)
	u.mediator.SendMessage(msg, u)
}

func (u *ChatUser) ReceiveMessage(msg string) {
	fmt.Printf("%s recibe: %s\n", u.name, msg)
}

func main() {
	chatRoom := &ChatRoom{}

	user1 := &ChatUser{name: "Juan", mediator: chatRoom}
	user2 := &ChatUser{name: "Maria", mediator: chatRoom}
	user3 := &ChatUser{name: "Carlos", mediator: chatRoom}

	chatRoom.AddUser(user1)
	chatRoom.AddUser(user2)
	chatRoom.AddUser(user3)

	user1.SendMessage("Hola a todos!")
}
```

---

## 5. Beneficios y Aplicaciones
### **Beneficios:**
✅ Reduce las dependencias entre objetos (desacoplamiento).
✅ Facilita la modificación y extensibilidad del sistema.
✅ Mejora la organización del código y la legibilidad.

### **Aplicaciones:**
- Sistemas de mensajería o chats.
- Controladores de interfaces gráficas.
- Coordinación entre módulos en software complejo.

---

## 6. Ejercicio Práctico
**Instrucciones:**
1. Implementa una versión del patrón Mediator en Go que simule el control del tráfico aéreo.
2. Define un mediador que coordine la comunicación entre diferentes aviones en una pista de aterrizaje.
3. Asegúrate de que los aviones soliciten permiso para aterrizar y despegar a través del mediador.
4. Prueba la implementación con varios aviones.

---

## 7. Recursos Adicionales
- [Refactoring Guru - Mediator Pattern](https://refactoring.guru/design-patterns/mediator)
- [Golang Design Patterns](https://github.com/tmrts/go-patterns)
- [Dive Into Design Patterns - Mediator](https://refactoring.guru/es/design-patterns/mediator)

---

**Fin de la Clase 18** 🎯

