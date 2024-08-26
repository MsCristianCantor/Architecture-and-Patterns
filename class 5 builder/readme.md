# Día 5: Patrón Builder

## Objetivo

Comprender y aplicar el **Patrón Builder** para construir objetos complejos paso a paso, permitiendo flexibilidad en su creación.

## Instrucciones

### 1. Concepto del Patrón Builder

El **Patrón Builder** es un patrón creacional que permite construir objetos complejos mediante la combinación de varios pasos sencillos. Este patrón es útil cuando un objeto puede ser construido de diferentes maneras o cuando su construcción implica varios pasos que deben seguirse en un orden específico.

### 2. Ventajas y Desventajas del Patrón Builder

- **Ventajas:**
  - Permite crear objetos paso a paso, mejorando la legibilidad del código.
  - Facilita la creación de diferentes representaciones de un objeto.
  - Separa la construcción de un objeto complejo de su representación final.
  
- **Desventajas:**
  - Puede introducir complejidad adicional si se usa en situaciones donde no es necesario.
  - Puede requerir la creación de múltiples clases adicionales, lo que aumenta la cantidad de código.

### 3. Implementación en Go

Crea un archivo `builder.go` y sigue estos pasos para implementar el patrón Builder en Go:

```go
package main

import "fmt"

// Producto final que queremos construir
type House struct {
    Foundation string
    Structure  string
    Roof       string
    Interior   string
}

// Builder es la interfaz que define los pasos para construir la casa
type Builder interface {
    SetFoundation()
    SetStructure()
    SetRoof()
    SetInterior()
    GetHouse() House
}

// ConcreteBuilder es una implementación concreta del Builder
type ConcreteBuilder struct {
    house House
}

func (b *ConcreteBuilder) SetFoundation() {
    b.house.Foundation = "Concrete, brick, and stone"
}

func (b *ConcreteBuilder) SetStructure() {
    b.house.Structure = "Wood and bricks"
}

func (b *ConcreteBuilder) SetRoof() {
    b.house.Roof = "Concrete and steel"
}

func (b *ConcreteBuilder) SetInterior() {
    b.house.Interior = "Wood and paint"
}

func (b *ConcreteBuilder) GetHouse() House {
    return b.house
}

// Director es responsable de manejar el proceso de construcción
type Director struct {
    builder Builder
}

func (d *Director) SetBuilder(b Builder) {
    d.builder = b
}

func (d *Director) ConstructHouse() {
    d.builder.SetFoundation()
    d.builder.SetStructure()
    d.builder.SetRoof()
    d.builder.SetInterior()
}

func main() {
    // Crear el Builder concreto y el Director
    builder := &ConcreteBuilder{}
    director := Director{builder: builder}

    // Construir la casa
    director.ConstructHouse()

    // Obtener el producto final
    house := builder.GetHouse()

    fmt.Printf("House built with: %s, %s, %s, %s\n",
        house.Foundation, house.Structure, house.Roof, house.Interior)
}
```

#### Explicación del Código

- Interfaz Builder: Define los métodos necesarios para construir las partes de un objeto complejo.
- Clase ConcreteBuilder: Implementa los métodos definidos en la interfaz Builder y mantiene el estado del objeto en construcción.
- Clase Director: Controla el proceso de construcción usando el Builder para crear el objeto paso a paso.
- Producto House: Representa el objeto complejo que se está construyendo.

### 4. Práctica

- Extiende el ConcreteBuilder para permitir la personalización del interior de la casa, como la elección de colores o materiales específicos.
- Crea otro Builder concreto que siga un proceso de construcción diferente, por ejemplo, para una casa de estilo moderno.
- Modifica el Director para que pueda construir diferentes tipos de casas usando distintos builders.

### 5. Reflexión

- Piensa en cómo el patrón Builder puede ayudar a manejar la creación de objetos complejos en sistemas grandes.
- Considera en qué situaciones podrías preferir el patrón Builder sobre otros patrones creacionales, como el Abstract Factory o el Factory Method.

### 6. Tareas Adicionales

- Investiga cómo se implementa el patrón Builder en otros lenguajes de programación.
- Compara el patrón Builder con el patrón Prototype. ¿En qué situaciones uno es preferible sobre el otro?

### 7. Recursos Adicionales

- Consulta la sección sobre el Builder en "Design Patterns: Elements of Reusable Object-Oriented Software" de Gamma et al.
- Explora ejemplos de Builder en la documentación de Go para ver cómo se utiliza en la práctica.

### Resumen

- El Patrón Builder es fundamental para la construcción de objetos complejos paso a paso, ofreciendo flexibilidad y claridad en el proceso de creación.
- Este patrón es útil cuando la creación de un objeto implica múltiples pasos que pueden variar en diferentes contextos.
