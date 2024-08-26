# Día 3: Patrón Factory Method

## Objetivo

Entender y aplicar el **Patrón Factory Method**, que permite la creación de objetos a través de una interfaz común sin especificar la clase exacta del objeto que será creado.

## Instrucciones

### 1. Concepto del Patrón Factory Method

El **Patrón Factory Method** es un patrón de diseño creacional que define una interfaz para crear un objeto, pero permite que las subclases alteren el tipo de objeto que se creará. Este patrón es útil cuando la creación del objeto implica lógica compleja o cuando no quieres que la lógica de creación esté dispersa por todo el código.

### 2. Ventajas y Desventajas del Factory Method

- **Ventajas:**
  - Fomenta el encapsulamiento de la lógica de creación de objetos.
  - Hace que el código sea más flexible y más fácil de mantener.
- **Desventajas:**
  - Puede añadir complejidad adicional si se abusa de su uso.

### 3. Implementación en Go

Crea un archivo `factory_method.go` y escribe el siguiente código para implementar el patrón Factory Method en Go:

```go
package main

import "fmt"

// Producto es la interfaz que define un comportamiento común
type Product interface {
    Use() string
}

// ConcreteProductA es una implementación concreta de Product
type ConcreteProductA struct{}

func (p ConcreteProductA) Use() string {
    return "Using Product A"
}

// ConcreteProductB es otra implementación concreta de Product
type ConcreteProductB struct{}

func (p ConcreteProductB) Use() string {
    return "Using Product B"
}

// Factory es la interfaz que define el método Factory
type Factory interface {
    CreateProduct() Product
}

// FactoryA es una implementación concreta de Factory que crea ProductA
type FactoryA struct{}

func (f FactoryA) CreateProduct() Product {
    return ConcreteProductA{}
}

// FactoryB es una implementación concreta de Factory que crea ProductB
type FactoryB struct{}

func (f FactoryB) CreateProduct() Product {
    return ConcreteProductB{}
}

func main() {
    var factory Factory

    // Crear un ProductA
    factory = FactoryA{}
    productA := factory.CreateProduct()
    fmt.Println(productA.Use())

    // Crear un ProductB
    factory = FactoryB{}
    productB := factory.CreateProduct()
    fmt.Println(productB.Use())
}
```

#### Explicación del Código

- Interfaz Product: Define el comportamiento que todos los productos deben tener.
- Clases ConcreteProductA y ConcreteProductB: Implementaciones específicas del producto.
- Interfaz Factory: Define un método para crear productos.
- Clases FactoryA y FactoryB: Implementaciones concretas del Factory Method, que crean productos específicos.

### 4. Práctica

- Modifica el ejemplo para añadir un tercer tipo de producto (por ejemplo, ConcreteProductC) y su correspondiente fábrica (FactoryC).
- Implementa un nuevo método en la interfaz Product que ambos productos deben cumplir, y asegúrate de que todas las fábricas lo implementan correctamente.

### 5. Reflexión

- Piensa en situaciones en las que podrías necesitar diferenciar la creación de objetos en función de condiciones específicas en tus proyectos.
- Considera cómo este patrón puede ayudarte a mantener tu código limpio y fácil de mantener al encapsular la lógica de creación.

### 6. Tareas Adicionales

- Investiga otros patrones de diseño relacionados como el Abstract Factory o Builder para comprender cómo podrías combinarlos con el Factory Method.
- Revisa ejemplos en la documentación de Go y en otros lenguajes para ver cómo se implementa el Factory Method en diferentes contextos.

### 7. Recursos Adicionales

- Lee más sobre el Factory Method en "Design Patterns: Elements of Reusable Object-Oriented Software" de Gamma et al.
- Explora artículos y videos sobre patrones creacionales en Go para obtener una comprensión más profunda.

### Resumen

- El Patrón Factory Method es útil para delegar la creación de objetos y mejorar la flexibilidad y mantenibilidad del código.
- Comprender este patrón te permitirá diseñar sistemas que sean fáciles de extender y modificar sin tocar la lógica existente.
