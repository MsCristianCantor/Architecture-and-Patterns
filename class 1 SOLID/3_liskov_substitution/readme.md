# Liskov Substitution

Tu interpretación del Principio de Sustitución de Liskov (Liskov Substitution Principle, LSP) es un buen punto de partida, pero aquí te dejo algunos ajustes y aclaraciones para hacerla más precisa:

* Herencia y Sustitución: El principio LSP va más allá de la simple herencia. Indica que si una clase B es una subclase de A, los objetos de la clase B deben poder sustituir a los objetos de la clase A sin alterar el comportamiento correcto del programa. En otras palabras, las clases derivadas deben ser completamente intercambiables con sus clases base sin introducir errores o comportamientos inesperados.

* Contrato: El LSP asegura que una subclase respete el "contrato" definido por su clase base. Esto significa que los métodos de la subclase deben cumplir con las expectativas establecidas por la clase base en términos de resultados, excepciones lanzadas, y cualquier precondición o postcondición.

## Ejemplo Mejorado

* "El principio de sustitución de Liskov establece que las subclases deben ser capaces de sustituir a sus clases base sin afectar la lógica del programa. Esto implica que la subclase debe cumplir con las expectativas y contratos definidos por la clase base, asegurando que no se introduzcan comportamientos inesperados o errores al reemplazar una clase por otra."

Este ajuste te ayudará a tener una comprensión más completa del LSP, que es crucial para mantener la coherencia y confiabilidad en sistemas que hacen uso de la herencia. ¡Continúa así!

## Explicación codigo

Aunque Penguin implementa Bird, no respeta el contrato esperado porque los pingüinos no vuelan, violando así el principio LSP.