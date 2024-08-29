# Clase 7: Adapter Pattern

## Objetivos de la clase

- Comprender el propósito y la estructura del patrón Adapter.
- Implementar un ejemplo del patrón Adapter en Go.
- Identificar cuándo es apropiado utilizar este patrón en la arquitectura de software.

## Descripción

El patrón Adapter se utiliza para permitir que clases con interfaces incompatibles trabajen juntas. Este patrón actúa como un puente entre dos interfaces, adaptando la interfaz de una clase existente para que coincida con la interfaz esperada por el cliente.

### Casos de uso comunes

- Integrar clases existentes en un sistema sin modificar su código.
- Adaptar una interfaz heredada a una nueva sin alterar la clase original.
- Facilitar la interacción entre componentes de software con interfaces diferentes.

## Instrucciones

### 1. Teoría del Adapter Pattern

- Revisa la definición y estructura del patrón Adapter en los materiales proporcionados.
- Entiende cómo este patrón permite la integración de clases con interfaces incompatibles.

### 2. Ejemplo en Go

A continuación, se presenta un ejemplo básico del patrón Adapter en Go:

```go
package main

import "fmt"

// Target es la interfaz que el cliente espera
type Target interface {
    Request() string
}

// Adaptee es una estructura que tiene una interfaz incompatible con Target
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
    return "Specific request"
}

// Adapter es la estructura que adapta la interfaz de Adaptee a Target
type Adapter struct {
    Adaptee *Adaptee
}

func (a *Adapter) Request() string {
    return a.Adaptee.SpecificRequest()
}

func main() {
    adaptee := &Adaptee{}
    adapter := &Adapter{Adaptee: adaptee}

    fmt.Println("Adapter result: ", adapter.Request())
}
```

### 3. Implementación práctica

- Implementa el ejemplo en tu entorno de desarrollo.
- Asegúrate de entender cómo el Adapter convierte la interfaz de Adaptee para que sea compatible con Target.

### 4. Ejercicio práctico

- Crea una situación donde tengas dos interfaces incompatibles en un proyecto.
- Utiliza el patrón Adapter para integrarlas sin modificar el código original de las clases.

### 5. Discusión y preguntas

- Reflexiona sobre cómo este patrón puede ser útil en proyectos reales.
- Prepárate para discutir tus preguntas y experiencias durante la próxima clase.

## Material adicional

Enlaces a recursos y documentación adicional sobre el Adapter Pattern.
