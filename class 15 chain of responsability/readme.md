# Clase 15: Chain of Responsibility Pattern

## 1. Objetivo de la clase
En esta clase, aprenderás sobre el patrón de diseño **Chain of Responsibility**, el cual permite procesar una solicitud a través de una cadena de handlers. Este patrón se utiliza para desacoplar el emisor de la solicitud de sus receptores.

Al finalizar la clase, podrás:
- Entender los principios del patrón Chain of Responsibility.
- Implementar este patrón en Go.
- Identificar casos prácticos donde este patrón es útil.

---

## 2. ¿Qué es el Chain of Responsibility?
El patrón Chain of Responsibility permite que múltiples objetos tengan la oportunidad de manejar una solicitud, eliminando la dependencia entre el emisor y el receptor de la misma. Este patrón se implementa encadenando objetos receptores y pasando la solicitud a lo largo de la cadena hasta que algún receptor la procese.

### Ventajas:
- Desacopla el emisor del receptor.
- Permite añadir o modificar handlers sin alterar el código existente.
- Mejora la flexibilidad al procesar solicitudes.

### Desventajas:
- Puede ser difícil de depurar debido a la naturaleza encadenada.
- No garantiza que una solicitud sea procesada.

---

## 3. Diagrama UML

![uml-chain-of-responsibility.jpg](uml-chain-of-responsibility.jpg)

---

## 4. Ejemplo conceptual
Supongamos que tienes un sistema de soporte técnico con distintos niveles de atención (soporte básico, técnico avanzado, y gerente). Cada solicitud debe ser atendida por el nivel correspondiente.

### Estructura en Go:
1. **Handler (Interfaz):** Define un método para procesar la solicitud y un método para pasarla al siguiente handler.
2. **ConcreteHandler:** Implementa la lógica específica para procesar la solicitud.
3. **Cliente:** Inicia la solicitud y la pasa al primer handler.

---

## 5. Ejercicio práctico
### Problema:
Implementa un sistema de validación de datos de entrada donde cada regla de validación sea un handler en la cadena.

#### Requisitos:
1. Si un dato no cumple con una regla, se debe detener el proceso de validación y mostrar el error correspondiente.
2. Las reglas a implementar son:
   - Verificar que la cadena no esté vacía.
   - Verificar que la cadena tenga al menos 5 caracteres.
   - Verificar que la cadena no contenga caracteres especiales.

#### Pasos:
1. Define la interfaz `Handler` con métodos para manejar solicitudes y configurar el siguiente handler.
2. Crea implementaciones concretas para cada regla de validación.
3. Crea una cadena de validación en el `main` y prueba con diferentes datos de entrada.

---

## 6. Recursos adicionales
- [Go Patterns: Chain of Responsibility](https://refactoring.guru/design-patterns/chain-of-responsibility/go/example)
- Video explicativo: [Chain of Responsibility Pattern en Go](https://www.youtube.com/watch?v=4UjM2D_TbTs)
- Libro: *Design Patterns: Elements of Reusable Object-Oriented Software* por Gamma et al.

---

## 7. Tarea
1. Implementa el ejercicio práctico propuesto.
2. Busca un caso real en tu trabajo o proyectos personales donde puedas aplicar el patrón Chain of Responsibility.
3. Documenta tus hallazgos y discútelos en la próxima clase.
