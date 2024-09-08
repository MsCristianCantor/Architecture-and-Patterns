# Clase 10: Composite Pattern

## Objetivo

Entender e implementar el **Patrón Composite** en Go, que permite a los clientes tratar objetos individuales y compuestos de manera uniforme. Este patrón es útil cuando los objetos forman parte de una jerarquía que puede representar tanto elementos simples como complejos.

## Conceptos Clave

- **Composite Pattern**: Este patrón permite organizar objetos en estructuras de árbol para representar jerarquías de "parte-todo". Los objetos compuestos pueden tener otros objetos dentro de ellos, y los clientes pueden interactuar con todos ellos de manera uniforme.
  
- **Component**: Define una interfaz para los objetos de la estructura.
  
- **Leaf**: Representa objetos individuales que no tienen hijos. Implementa los métodos de la interfaz `Component`.

- **Composite**: Un objeto que puede tener otros componentes (hojas u otros composites) y que también implementa la interfaz `Component`.

## Instrucciones

### 1. Crear la Interfaz `Component`

Define una interfaz que declare las operaciones comunes que todos los elementos (simples y compuestos) deben implementar.

```go
type Component interface {
    Operation() string
    Add(Component)
    Remove(Component)
    GetChild(int) Component
}
```

### 2. Crear la Clase `Leaf`

Implementa la interfaz `Component` para representar objetos simples.

```go
type Leaf struct {
    name string
}

func (l *Leaf) Operation() string {
    return l.name
}

func (l *Leaf) Add(c Component) {
    // No hace nada ya que es una hoja
}

func (l *Leaf) Remove(c Component) {
    // No hace nada ya que es una hoja
}

func (l *Leaf) GetChild(i int) Component {
    return nil // Las hojas no tienen hijos
}
```

### 3. Crear la Clase `Composite`

Implementa la interfaz `Component` para objetos compuestos que pueden contener otros componentes.

```go
type Composite struct {
    children []Component
}

func (c *Composite) Operation() string {
    result := "Composite ["
    for _, child := range c.children {
        result += child.Operation() + " "
    }
    result += "]"
    return result
}

func (c *Composite) Add(component Component) {
    c.children = append(c.children, component)
}

func (c *Composite) Remove(component Component) {
    // Implementar lógica para eliminar el componente
}

func (c *Composite) GetChild(i int) Component {
    if i >= 0 && i < len(c.children) {
        return c.children[i]
    }
    return nil
}
```

### 4. Implementar el Código Cliente

Crea un objeto compuesto que contenga tanto objetos simples como otros compuestos.

```go
func main() {
    leaf1 := &Leaf{name: "Leaf 1"}
    leaf2 := &Leaf{name: "Leaf 2"}
    composite := &Composite{}
    composite.Add(leaf1)
    composite.Add(leaf2)

    fmt.Println(composite.Operation()) // Salida: Composite [Leaf 1 Leaf 2 ]
}
```

### 5. Ejercicio Práctico

1. Implementa un sistema de archivos donde:

    - Los archivos son las hojas.
    - Las carpetas son compuestos que pueden contener tanto archivos como otras carpetas.

2. Crea una jerarquía de carpetas con archivos y carpetas dentro, y realiza una operación para listar todos los nombres de los archivos y carpetas.

#### Recursos Adicionales

- [Composite Pattern en Go](https://refactoring.guru/design-patterns/composite/go/example)
- [Documentación oficial de Go](https://golang.org/doc/)
