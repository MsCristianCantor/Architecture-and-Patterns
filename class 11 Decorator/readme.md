
# Clase 11: Decorator Pattern

## Objetivo

Al finalizar esta clase, los estudiantes comprenderán el uso del patrón de diseño **Decorator** y cómo puede ser implementado para agregar dinámicamente funcionalidades adicionales a objetos sin modificar su estructura.

## Contenido

1. **Introducción al patrón Decorator**
2. **Diagrama UML del patrón Decorator**
3. **Componentes del patrón Decorator**
4. **Ventajas y desventajas**
5. **Caso práctico con código en Go**

---

### 1. Introducción al patrón Decorator

El **Decorator Pattern** es un patrón estructural que permite agregar funcionalidades a un objeto de manera dinámica sin alterar su código. El patrón es útil cuando queremos extender las funcionalidades de una clase, pero sin heredar de ella ni modificarla directamente.

### 2. Diagrama UML del patrón Decorator

![UML Decorator Pattern](uml-decorator-pattern.png)

En el diagrama:

- **Component**: Define la interfaz base para los objetos que pueden tener funcionalidades agregadas dinámicamente.
- **ConcreteComponent**: Implementa la interfaz `Component` y es el objeto principal que puede recibir decoradores.
- **Decorator**: Mantiene una referencia al `Component` y define una interfaz que sigue siendo compatible con la interfaz base.
- **ConcreteDecorator**: Extiende las funcionalidades del objeto decorado.

### 3. Componentes del patrón Decorator

1. **Componente (Component)**: Interfaz o clase abstracta que define las operaciones que pueden ser decoradas.
2. **Componente Concreto (ConcreteComponent)**: La implementación principal de la interfaz `Component`.
3. **Decorador (Decorator)**: Clase abstracta que implementa la interfaz `Component` y contiene una referencia a un `Component`.
4. **Decorador Concreto (ConcreteDecorator)**: Implementación concreta que agrega una nueva funcionalidad al `ConcreteComponent`.

### 4. Ventajas y desventajas

#### Ventajas

- **Flexibilidad**: Permite extender las funcionalidades de manera flexible y dinámica.
- **Abierto para extensión, cerrado para modificación**: Siguiendo el principio SOLID.

#### Desventajas

- **Complejidad**: Puede hacer el código más complicado de entender debido al número de clases adicionales que se crean.
- **Orden de los decoradores**: El orden en el que se aplican los decoradores puede afectar el resultado.

### 5. Caso práctico con código en Go

```go
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
```

---

## Ejercicio práctico

### Instrucciones

1. Implementa un sistema de notificaciones en el que puedas decorar mensajes con diferentes prioridades (baja, media, alta).
2. Crea una clase base `Notificador` que defina el mensaje principal.
3. Agrega decoradores que permitan adjuntar etiquetas como [Urgente] o [Confidencial] al mensaje.

---

## Recursos adicionales

- [Refactoring.Guru: Decorator Pattern](https://refactoring.guru/design-patterns/decorator)
- [GoF Design Patterns](https://en.wikipedia.org/wiki/Decorator_pattern)
