# Día 2: Patrón Singleton

## Objetivo

Entender el **Patrón Singleton** y cómo implementarlo en Go para asegurar que una clase tenga solo una instancia y proporcione un punto global de acceso a esa instancia.

## Instrucciones

### 1. Concepto del Patrón Singleton

El Patrón Singleton asegura que una clase tenga solo una instancia y proporciona un punto de acceso global a dicha instancia. Esto es útil en situaciones donde necesitas un único objeto para coordinar acciones en todo el sistema, como en el caso de gestores de conexiones de bases de datos, registros de logs, o configuraciones globales.

### 2. Ventajas y Desventajas del Singleton

- **Ventajas:**
  - Controla el acceso a una única instancia.
  - Reduce el uso de recursos y asegura un comportamiento uniforme.
- **Desventajas:**
  - Puede dificultar las pruebas unitarias si no se diseña correctamente.
  - Introduce un acoplamiento más fuerte en el código.

### 3. Implementación en Go

Crea un archivo `singleton.go` y escribe el siguiente código para implementar el patrón Singleton en Go:

```go
package main

import (
    "fmt"
    "sync"
)

// Singleton estructura que queremos asegurar que solo tenga una instancia
type Singleton struct {
    name string
}

var instance *Singleton
var once sync.Once

// GetInstance retorna la instancia única del Singleton
func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{name: "UniqueInstance"}
    })
    return instance
}

func main() {
    singleton1 := GetInstance()
    singleton2 := GetInstance()

    if singleton1 == singleton2 {
        fmt.Println("Both variables point to the same instance.")
    }
}
```

#### Explicación del Código

- Usamos sync.Once para asegurar que la inicialización de la instancia ocurra solo una vez, incluso si hay concurrencia.
- GetInstance() es el punto global de acceso que garantiza la única instancia del objeto.

### 4. Práctica

- Modifica el ejemplo para que la instancia de Singleton almacene alguna configuración global y prueba su uso en diferentes partes del código.
- Asegúrate de entender cómo funciona el control de concurrencia con sync.Once para garantizar que la instancia se crea solo una vez, incluso en entornos concurrentes.

### 5. Reflexión

- Piensa en situaciones en las que podrías necesitar una única instancia de un objeto en tus proyectos. ¿Dónde podrías aplicar el Patrón Singleton?
- Considera los pros y contras del uso del Singleton en tu arquitectura y cómo puede afectar la escalabilidad y testabilidad del sistema.

### 6. Tareas Adicionales

- Investiga las alternativas al Singleton en Go y cómo manejar escenarios donde el Singleton podría crear problemas, como en sistemas distribuidos o en contextos donde la testabilidad es clave.

### 7. Recursos Adicionales

- Revisa más sobre el Patrón Singleton en el libro "Design Patterns: Elements of Reusable Object-Oriented Software" de Gamma et al.
- Explora cómo otros lenguajes de programación implementan Singleton y compara con la implementación en Go.

### Resumen

- El Patrón Singleton es fundamental para asegurar que solo exista una instancia de una clase en todo el sistema.
- Es importante entender cuándo y cómo utilizarlo adecuadamente para evitar complicaciones en el código.
