# Día 6: Patrón Prototype

## Objetivo

Comprender y aplicar el **Patrón Prototype** para crear nuevos objetos copiando una instancia existente en lugar de crearla desde cero.

## Instrucciones

### 1. Concepto del Patrón Prototype

El **Patrón Prototype** es un patrón creacional que permite crear nuevos objetos clonando una instancia existente. Este patrón es útil cuando la creación de un objeto es costosa en términos de tiempo o recursos, y se puede crear un clon de un objeto ya existente para obtener un nuevo objeto con las mismas características.

### 2. Ventajas y Desventajas del Patrón Prototype

- **Ventajas:**
  - Reduce la necesidad de subclases adicionales al permitir la clonación de objetos.
  - Mejora el rendimiento cuando la clonación es más rápida que la creación de un objeto desde cero.
  - Facilita la creación de objetos complejos que ya están configurados.

- **Desventajas:**
  - Puede ser complejo implementar un clon profundo si los objetos contienen referencias a otros objetos.
  - No es siempre evidente qué objetos deberían ser clonables.

### 3. Implementación en Go

Crea un archivo `prototype.go` y sigue estos pasos para implementar el patrón Prototype en Go:

```go
package main

import "fmt"

// Clonable es la interfaz que define el método Clone
type Clonable interface {
    Clone() Clonable
}

// ConcretePrototype es una implementación concreta de Clonable
type ConcretePrototype struct {
    Name  string
    Value int
}

func (p *ConcretePrototype) Clone() Clonable {
    clone := *p
    return &clone
}

func main() {
    // Crear un objeto prototype
    prototype := &ConcretePrototype{Name: "Prototype1", Value: 100}

    // Clonar el prototype
    clone := prototype.Clone().(*ConcretePrototype)
    clone.Name = "Prototype2"

    // Mostrar los resultados
    fmt.Printf("Original Prototype: %s, %d\n", prototype.Name, prototype.Value)
    fmt.Printf("Cloned Prototype: %s, %d\n", clone.Name, clone.Value)
}
```

### 4. Práctica

- Extiende el ConcretePrototype para incluir un campo adicional como una lista o mapa, y modifica el método Clone para manejar la clonación profunda de estos campos.
- Crea diferentes tipos de prototypes que compartan alguna estructura común pero que varíen en algunos de sus campos.

### 5. Reflexión

- Considera en qué situaciones el patrón Prototype sería más ventajoso que el uso directo de constructores o fábricas.
- Piensa en los posibles problemas que podrías enfrentar al implementar la clonación profunda, especialmente en objetos complejos o con ciclos de referencias.

### 6. Tareas Adicionales

- Investiga cómo se implementa la clonación profunda en otros lenguajes de programación y compáralo con Go.
- Explora cómo el patrón Prototype puede integrarse con otros patrones creacionales, como el Builder o el Factory Method.

### 7. Recursos Adicionales

- Consulta la sección sobre Prototype en "Design Patterns: Elements of Reusable Object-Oriented Software" de Gamma et al.
- Revisa ejemplos de clonación en la documentación de Go para ver cómo se utiliza en la práctica.

### Resumen

- El Patrón Prototype es útil para crear nuevos objetos mediante la clonación de instancias existentes, lo que puede ahorrar tiempo y recursos en la creación de objetos complejos.
-Este patrón es especialmente útil cuando la creación de un objeto es costosa o cuando se necesita preservar el estado de un objeto en un nuevo objeto.