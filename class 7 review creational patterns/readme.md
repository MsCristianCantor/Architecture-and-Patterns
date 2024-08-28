# Clase 7: Revisión de Patrones Creacionales y Ejercicios Prácticos

## Objetivo de la Clase

En esta clase, revisaremos todos los patrones de diseño creacionales estudiados hasta ahora y realizaremos ejercicios prácticos para consolidar los conocimientos adquiridos.

## Resumen de Patrones Creacionales Estudiados

1. **Singleton**
   - Asegura que una clase solo tenga una instancia y proporciona un punto global de acceso a ella.
   - Ejemplo en Go: Implementación de un singleton con inicialización perezosa utilizando `sync.Once`.

2. **Factory Method**
   - Define una interfaz para crear objetos en una superclase, pero permite que las subclases alteren el tipo de objetos que se crearán.
   - Ejemplo en Go: Crear diferentes productos a través de fábricas específicas.

3. **Abstract Factory**
   - Proporciona una interfaz para crear familias de objetos relacionados o dependientes sin especificar sus clases concretas.
   - Ejemplo en Go: Implementación de diferentes familias de muebles, como sillas y sofás de estilo victoriano, moderno y art déco.

4. **Builder**
   - Separa la construcción de un objeto complejo de su representación, de manera que el mismo proceso de construcción puede crear diferentes representaciones.
   - Ejemplo en Go: Construcción de casas con diferentes características utilizando el patrón Builder.

5. **Prototype**
   - Permite copiar objetos existentes sin que el código dependa de sus clases.
   - Ejemplo en Go: Clonar objetos que implementan la interfaz `Clonable`.

## Actividades

### Actividad 1: Revisión Teórica

- **Paso 1:** Repasa los conceptos clave de cada uno de los patrones creacionales estudiados.
- **Paso 2:** Compara las diferencias y similitudes entre los patrones. Reflexiona sobre cuándo es más apropiado usar cada uno.

### Actividad 2: Implementación Práctica

- **Paso 1:** Elige uno de los patrones creacionales que has estudiado.
- **Paso 2:** Piensa en un problema o caso de uso en el que este patrón pueda ser aplicado.
- **Paso 3:** Implementa una solución en Go utilizando el patrón elegido.
- **Paso 4:** Comparte tu implementación con tus compañeros o revisa las soluciones de otros para identificar mejoras o diferentes enfoques.

### Actividad 3: Ejercicio Comparativo

- **Paso 1:** Toma un problema de diseño simple (por ejemplo, la creación de un vehículo).
- **Paso 2:** Implementa soluciones utilizando diferentes patrones creacionales (Factory Method, Abstract Factory, Builder).
- **Paso 3:** Discute cuál de las implementaciones es más flexible, más fácil de mantener, y por qué.

## Reflexión Final

- ¿Cuál de los patrones creacionales te parece más útil en tu contexto actual de trabajo?
- ¿Cómo podrías aplicar estos patrones en proyectos futuros?

## Tarea

- **Tarea:** Realiza una implementación completa utilizando el patrón Builder o Abstract Factory para un proyecto personal o un caso de estudio propuesto. Prepárate para presentar y discutir tu solución en la próxima clase.

## Recursos Adicionales

- [GoF Design Patterns](https://refactoring.guru/design-patterns/go)
- [Golang Patterns](https://golangpatterns.info/)

---

**Fecha de la Clase:** `[Fecha de la clase]`
