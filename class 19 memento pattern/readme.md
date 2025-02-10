# Clase 19: Memento Pattern

## 1. Introducción
El **Memento Pattern** es un patrón de diseño de comportamiento que permite capturar y restaurar el estado de un objeto sin violar su encapsulamiento. Es útil cuando se requiere implementar funciones de "deshacer" o "rehacer" en una aplicación.

## 2. Objetivo de Aprendizaje
Al finalizar esta clase, deberás ser capaz de:
- Comprender el concepto y utilidad del **Memento Pattern**.
- Implementar el patrón en Go.
- Aplicarlo en escenarios del mundo real como editores de texto o juegos.

## 3. Conceptos Clave
- **Memento**: Representa un snapshot del estado del objeto.
- **Originator**: Crea y usa los Mementos para restaurar su estado.
- **Caretaker**: Almacena y administra los Mementos.

## 4. Diagrama UML
(Asegúrate de revisar el diagrama UML incluido en los recursos adicionales para visualizar la estructura del patrón.)

## 5. Implementación en Go
A continuación, se presenta una implementación básica del **Memento Pattern** en Go:

```go
package main

import "fmt"

// Memento almacena el estado de un objeto
 type Memento struct {
    state string
}

// Originator crea y usa Mementos
 type Originator struct {
    state string
}

func (o *Originator) Save() *Memento {
    return &Memento{state: o.state}
}

func (o *Originator) Restore(m *Memento) {
    o.state = m.state
}

// Caretaker administra los Mementos
 type Caretaker struct {
    history []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
    c.history = append(c.history, m)
}

func (c *Caretaker) GetMemento(index int) *Memento {
    if index < len(c.history) {
        return c.history[index]
    }
    return nil
}

func main() {
    originator := &Originator{state: "Estado Inicial"}
    caretaker := &Caretaker{}

    // Guardar estado
    caretaker.AddMemento(originator.Save())
    originator.state = "Estado Modificado"
    fmt.Println("Estado actual:", originator.state)

    // Restaurar estado anterior
    originator.Restore(caretaker.GetMemento(0))
    fmt.Println("Estado restaurado:", originator.state)
}
```

## 6. Ejercicio Práctico
Implementa una versión del **Memento Pattern** en Go donde un usuario pueda escribir texto y deshacer cambios.

### Pasos:
1. Crea una estructura `Editor` con una propiedad `content`.
2. Implementa el método `Save()` que devuelve un Memento con el estado actual del contenido.
3. Implementa el método `Restore(memento)` que restablece el estado guardado.
4. Usa una estructura `History` para almacenar Mementos y permitir deshacer cambios.
5. Escribe una función `main()` que pruebe la funcionalidad.

## 7. Recursos Adicionales
- [Explicación del Memento Pattern en Refactoring Guru](https://refactoring.guru/design-patterns/memento)
- [Patrones de Diseño en Golang](https://golangbyexample.com/)
- **Diagrama UML** del patrón Memento (archivo "uml-memento-pattern.png")

---

Con esta clase, ahora tienes el conocimiento para aplicar el **Memento Pattern** en tus proyectos. ¡Manos a la obra! 🚀

