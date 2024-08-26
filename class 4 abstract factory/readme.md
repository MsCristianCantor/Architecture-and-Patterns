# Día 4: Patrón Abstract Factory

## Objetivo

Entender y aplicar el **Patrón Abstract Factory** para crear familias de objetos relacionados o dependientes sin especificar sus clases concretas.

## Instrucciones

### 1. Concepto del Patrón Abstract Factory

El **Patrón Abstract Factory** es un patrón creacional que proporciona una interfaz para crear familias de objetos relacionados o dependientes sin especificar sus clases concretas. Este patrón es útil cuando se quiere asegurarse de que los productos creados son compatibles entre sí.

### 2. Ventajas y Desventajas del Patrón Abstract Factory

- **Ventajas:**
  - Garantiza la compatibilidad entre los productos creados por la misma fábrica.
  - Facilita la sustitución de familias completas de productos sin cambiar el código cliente.
  - Promueve la coherencia en la creación de objetos relacionados.
  
- **Desventajas:**
  - Añade complejidad al código debido a la necesidad de crear varias clases adicionales.
  - Puede ser difícil de implementar cuando las familias de productos requieren muchas variaciones.

### 3. Implementación en Go

Crea un archivo `abstract_factory.go` y sigue estos pasos para implementar el patrón Abstract Factory en Go:

```go
package main

import "fmt"

// Interfaces para los productos abstractos
type Chair interface {
    SitOn() string
}

type Sofa interface {
    LayOn() string
}

// Productos concretos de la familia 1
type VictorianChair struct{}

func (v VictorianChair) SitOn() string {
    return "Sitting on a Victorian Chair"
}

type VictorianSofa struct{}

func (v VictorianSofa) LayOn() string {
    return "Laying on a Victorian Sofa"
}

// Productos concretos de la familia 2
type ModernChair struct{}

func (m ModernChair) SitOn() string {
    return "Sitting on a Modern Chair"
}

type ModernSofa struct{}

func (m ModernSofa) LayOn() string {
    return "Laying on a Modern Sofa"
}

// Abstract Factory
type FurnitureFactory interface {
    CreateChair() Chair
    CreateSofa() Sofa
}

// Fábrica concreta de la familia 1
type VictorianFurnitureFactory struct{}

func (v VictorianFurnitureFactory) CreateChair() Chair {
    return VictorianChair{}
}

func (v VictorianFurnitureFactory) CreateSofa() Sofa {
    return VictorianSofa{}
}

// Fábrica concreta de la familia 2
type ModernFurnitureFactory struct{}

func (m ModernFurnitureFactory) CreateChair() Chair {
    return ModernChair{}
}

func (m ModernFurnitureFactory) CreateSofa() Sofa {
    return ModernSofa{}
}

func main() {
    // Cliente usando la fábrica Victoriana
    var factory FurnitureFactory
    factory = VictorianFurnitureFactory{}

    chair := factory.CreateChair()
    sofa := factory.CreateSofa()

    fmt.Println(chair.SitOn())
    fmt.Println(sofa.LayOn())

    // Cliente usando la fábrica Moderna
    factory = ModernFurnitureFactory{}

    chair = factory.CreateChair()
    sofa = factory.CreateSofa()

    fmt.Println(chair.SitOn())
    fmt.Println(sofa.LayOn())
}
```

#### Explicación del Código

- Interfaces Chair y Sofa: Definen los métodos que los productos de las distintas familias deben implementar.
- Clases concretas (VictorianChair, ModernChair, etc.): Implementan las interfaces y representan los productos concretos de cada familia.
- Fábricas concretas (VictorianFurnitureFactory, ModernFurnitureFactory): Crean los productos concretos para cada familia.
- Cliente: Usa la fábrica abstracta para crear productos, sin preocuparse por las clases concretas.

### 4. Práctica

- Añade una nueva familia de productos, por ejemplo, muebles Art Deco, implementando las clases correspondientes para Chair y Sofa.
- Crea una nueva fábrica concreta para la familia Art Deco.
- Modifica el código del cliente para que use la nueva fábrica y pruebe los nuevos productos.

### 5. Reflexión

- Piensa en cómo este patrón podría ser útil en la construcción de interfaces gráficas, donde diferentes temas (familias de productos) podrían ser aplicados sin cambiar la lógica de la aplicación.
- Reflexiona sobre los posibles desafíos de mantener y extender el código cuando se añaden más familias de productos.

### 6. Tareas Adicionales

- Investiga cómo se implementa el patrón Abstract Factory en otros lenguajes de programación.
- Compara el patrón Abstract Factory con el patrón Factory Method. ¿En qué situaciones uno es preferible sobre el otro?

### 7. Recursos Adicionales

- Consulta la sección sobre el Abstract Factory en "Design Patterns: Elements of Reusable Object-Oriented Software" de Gamma et al.
- Revisa ejemplos de Abstract Factory en la documentación de Go o en proyectos open-source para ver cómo se utiliza en la práctica.

### Resumen

- El Patrón Abstract Factory es clave para la creación de familias de objetos relacionados o dependientes de manera consistente y extensible.
- Este patrón permite cambiar fácilmente entre familias de productos, asegurando la compatibilidad entre los objetos creados.
