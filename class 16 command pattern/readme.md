# Clase 16: Command Pattern

## 1. Objetivos de la clase
- Comprender el propósito y la estructura del patrón de diseño **Command**.
- Implementar el patrón **Command** en un ejemplo práctico usando Go.
- Aprender a encapsular solicitudes como objetos, lo que permite parametrizar invocaciones y soportar operaciones como "deshacer" (undo).

---

## 2. ¿Qué es el Command Pattern?

El patrón **Command** es un patrón de diseño de comportamiento que convierte una solicitud en un objeto independiente que contiene toda la información sobre la solicitud. Esto permite desacoplar al objeto que envía la solicitud del objeto que la procesa.

---

## 3. Componentes principales del Command Pattern

1. **Command (Interfaz):** Define un método para ejecutar una operación.
2. **ConcreteCommand:** Implementa la interfaz `Command` y encapsula una acción.
3. **Invoker:** Almacena y ejecuta los comandos.
4. **Receiver:** El objeto que realiza la acción real cuando se ejecuta el comando.

---

## 4. Ventajas del Command Pattern
- **Desacoplamiento:** Separa el emisor del receptor.
- **Extensibilidad:** Permite agregar nuevos comandos sin modificar el código existente.
- **Soporte para "undo" y "redo":** Fácil de implementar al almacenar el estado en los comandos.

---

## 5. Ejemplo práctico en Go

### Código inicial:
Vamos a implementar un sistema de control remoto donde los comandos controlan dispositivos como luces y ventiladores.

```go
package main

import "fmt"

// Command Interface
type Command interface {
	Execute()
	Undo()
}

// Receiver: Light
type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("Luz encendida")
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("Luz apagada")
}

// ConcreteCommand: LightOnCommand
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

func (c *LightOnCommand) Undo() {
	c.light.TurnOff()
}

// ConcreteCommand: LightOffCommand
type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}

func (c *LightOffCommand) Undo() {
	c.light.TurnOn()
}

// Invoker: RemoteControl
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func (r *RemoteControl) PressUndo() {
	r.command.Undo()
}

func main() {
	// Receiver
	light := &Light{}

	// Commands
	lightOn := &LightOnCommand{light: light}
	lightOff := &LightOffCommand{light: light}

	// Invoker
	remote := &RemoteControl{}

	// Encender la luz
	remote.SetCommand(lightOn)
	remote.PressButton()
	remote.PressUndo()

	// Apagar la luz
	remote.SetCommand(lightOff)
	remote.PressButton()
	remote.PressUndo()
}
```
Salida esperada:

    Luz encendida
    Luz apagada
    Luz apagada
    Luz encendida

---

## 6. Ejercicio práctico

### Enunciado:

Implementa un sistema de automatización de hogar donde puedas controlar múltiples dispositivos:

1. Un sistema de música con comandos para reproducir, pausar y detener.
2. Un sistema de ventilación con comandos para encender, apagar y ajustar la velocidad.

### Requisitos:

* Crea una interfaz Command para encapsular las acciones.
* Diseña un Invoker (controlador) que pueda manejar múltiples comandos.
* Implementa los comandos específicos y sus receptores.

---

## 7. Recursos adicionales

Aquí tienes algunos recursos útiles para profundizar en el patrón **Command**:

- [Refactoring Guru - Command Pattern](https://refactoring.guru/design-patterns/command): Explicación detallada del patrón con ejemplos en varios lenguajes de programación.
- [Go Patterns - Command](https://github.com/tmrts/go-patterns/blob/master/behavioral/command.md): Ejemplos y explicación del patrón Command en Go.
- [Design Patterns Explained - Command Pattern (YouTube)](https://www.youtube.com): Un video explicativo que muestra cómo implementar el patrón Command paso a paso.
- [Programación avanzada en Go](https://golang.org/doc/): La documentación oficial de Go siempre es un buen recurso para aprender más sobre programación avanzada.
- Libro recomendado: *"Head First Design Patterns"* - Un enfoque visual y práctico para entender patrones de diseño, incluido el Command.

¡Explora estos recursos para fortalecer tu comprensión! 🚀
