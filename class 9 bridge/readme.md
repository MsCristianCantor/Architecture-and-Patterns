# Clase 9: Bridge Pattern

## Objetivo

En esta clase, aprenderás a implementar el patrón de diseño **Bridge** en Go. Este patrón es útil para desacoplar una abstracción de su implementación, permitiendo que ambas varíen de manera independiente. Es especialmente útil cuando una clase se enfrenta a cambios en múltiples dimensiones o tiene múltiples variaciones posibles.

## Contenido

### 1. Teoría del Bridge Pattern

- **Definición**: El Bridge Pattern separa una abstracción de su implementación, permitiendo que ambas evolucionen de forma independiente. Es útil en escenarios donde se quiere evitar una explosión de clases derivadas que combinan múltiples variantes.
- **Estructura**:
  - **Abstracción**: Define la interfaz común para las abstracciones.
  - **Implementor**: Define la interfaz para las implementaciones.
  - **Concreta Abstracción**: Implementa la interfaz de la abstracción, y está compuesta de un `Implementor`.
  - **Concreto Implementor**: Implementa la interfaz `Implementor`.

### 2. Ejemplos Prácticos

- **Escenario**: Considera que estás diseñando un sistema de dispositivos electrónicos con diferentes modos de encendido (p. ej., modo remoto, manual).
  - La abstracción puede ser un `Device` (p. ej., `TV`, `Radio`).
  - La implementación puede ser un `RemoteControl` (p. ej., `BasicRemote`, `AdvancedRemote`).

### 3. Implementación en Go

1. **Define la interfaz de la Abstracción**: Crea una interfaz para `Device` que tenga métodos como `PowerOn` y `PowerOff`.

2. **Define la interfaz de la Implementación**: Crea una interfaz para `RemoteControl` con métodos como `TurnOn` y `TurnOff`.

3. **Crea las Clases Concretas para la Abstracción**: Implementa clases concretas como `TV` y `Radio` que cumplan la interfaz `Device`.

4. **Crea las Clases Concretas para la Implementación**: Implementa clases como `BasicRemote` y `AdvancedRemote` que cumplan la interfaz `RemoteControl`.

5. **Integración del Bridge**: Crea un puente entre el `Device` y el `RemoteControl`, de modo que puedas utilizar ambos de manera intercambiable.

### 4. Ejercicio Práctico

- **Instrucciones**:
  1. Implementa el patrón Bridge para un sistema de dibujo. Considera que hay dos tipos de formas (`Circle`, `Square`) y dos formas de dibujarlas (`DrawAPI1`, `DrawAPI2`).
  2. Define la abstracción para las formas y la implementación para los métodos de dibujo.
  3. Implementa al menos una clase concreta para cada uno.

- **Punto 4**:
  - Crea una instancia de cada combinación de forma y método de dibujo.
  - Llama a los métodos de dibujo para demostrar que la abstracción y la implementación están separadas.

### 5. Revisión de Código

- **Pasos**:
  - Envía tu implementación para revisión.
  - Discute posibles mejoras y variaciones del patrón.

## Recursos Adicionales

- [Documentación oficial de Go](https://golang.org/doc/)
- Ejemplos en [Refactoring.Guru](https://refactoring.guru/design-patterns/bridge)
- Artículos y tutoriales sobre Bridge Pattern en Go.

## Conclusión

Al finalizar la clase, deberás entender cómo implementar el Bridge Pattern y cuándo aplicarlo en sistemas que requieren una separación clara entre la abstracción y su implementación.
