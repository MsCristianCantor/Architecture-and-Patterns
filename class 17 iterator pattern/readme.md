# Clase 17: Iterator Pattern

## 1. Objetivo de la Clase
Comprender el **patrón de diseño Iterator** y su aplicación en Go para proporcionar una manera uniforme de recorrer colecciones de datos sin exponer su estructura interna.

## 2. Introducción al Patrón Iterator
El **Iterator Pattern** es un patrón de comportamiento que permite recorrer una colección de elementos secuencialmente sin exponer su implementación interna. Se basa en los siguientes componentes:

- **Iterador (Iterator)**: Interfaz que define los métodos para recorrer una colección.
- **Iterador Concreto (Concrete Iterator)**: Implementa la interfaz del iterador y mantiene el estado de la iteración.
- **Colección (Aggregate)**: Interfaz que define la estructura de datos iterable.
- **Colección Concreta (Concrete Aggregate)**: Implementa la interfaz de la colección y devuelve una instancia de un iterador.

## 3. Implementación en Go
Go no tiene iteradores incorporados como otros lenguajes, pero podemos implementar el patrón mediante interfaces y estructuras.

### Paso 1: Definir la interfaz Iterator
```go
// Iterator define los métodos de iteración
 type Iterator interface {
     HasNext() bool
     Next() interface{}
 }
```

### Paso 2: Definir la estructura de datos iterable
```go
// Collection define la interfaz para obtener un iterador
 type Collection interface {
     CreateIterator() Iterator
 }
```

### Paso 3: Implementar la colección concreta
```go
type ConcreteCollection struct {
    items []string
}

func (c *ConcreteCollection) CreateIterator() Iterator {
    return &ConcreteIterator{
        collection: c,
        index:      0,
    }
}
```

### Paso 4: Implementar el iterador concreto
```go
type ConcreteIterator struct {
    collection *ConcreteCollection
    index      int
}

func (i *ConcreteIterator) HasNext() bool {
    return i.index < len(i.collection.items)
}

func (i *ConcreteIterator) Next() interface{} {
    if i.HasNext() {
        item := i.collection.items[i.index]
        i.index++
        return item
    }
    return nil
}
```

### Paso 5: Usar el patrón en el código cliente
```go
func main() {
    collection := &ConcreteCollection{
        items: []string{"A", "B", "C", "D"},
    }
    iterator := collection.CreateIterator()

    for iterator.HasNext() {
        fmt.Println(iterator.Next())
    }
}
```

## 4. Beneficios del Patrón Iterator
- **Encapsulación**: No expone la estructura interna de la colección.
- **Flexibilidad**: Se pueden definir diferentes estrategias de recorrido sin modificar la colección.
- **Consistencia**: Ofrece una manera uniforme de recorrer diferentes tipos de colecciones.

## 5. Casos de Uso
- Recorrer listas, mapas o estructuras personalizadas sin exponer su implementación interna.
- Implementar diferentes estrategias de iteración, como iteradores inversos o filtrados.

## 6. Ejercicio Práctico
Implementa una estructura **Lista de Usuarios** que almacene nombres y proporcione un iterador para recorrerlos secuencialmente.

## 7. Recursos Adicionales
- [Patrón Iterator en Go](https://refactoring.guru/design-patterns/iterator/go/example)
- [Iteradores en Go](https://golangdocs.com/iterators-in-golang)

