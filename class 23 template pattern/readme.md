# Clase 23: Template Method Pattern

## 🧠 Objetivos

- Comprender el propósito y la estructura del patrón Template Method.
- Implementar el patrón en Go mediante ejemplos prácticos.
- Identificar casos reales de uso y cómo aplicar este patrón para mejorar la reutilización de código.
- Practicar la extensión de clases a través de métodos definidos por una estructura base.

---

## 📌 ¿Qué es el Patrón Template Method?

El patrón **Template Method** define el esqueleto de un algoritmo en una clase base, dejando que las subclases redefinan ciertos pasos sin cambiar la estructura general del algoritmo.

Este patrón pertenece a los **patrones de comportamiento** y se usa comúnmente para **evitar la duplicación de código** en algoritmos que comparten una estructura similar.

---

## 📊 Diagrama UML

![UML Template Pattern](https://refactoring.guru/images/patterns/diagrams/template-method/structure.png)

- `AbstractClass`: define el template method y los pasos del algoritmo.
- `ConcreteClass`: implementa los pasos específicos.

Fuente: [Refactoring Guru](https://refactoring.guru/design-patterns/template-method)

---

## 🧪 Ejemplo Conceptual

Supongamos que queremos definir el proceso de "hacer una bebida caliente", con pasos comunes como "hervir agua" y "verter en taza", pero con pasos específicos según sea té o café.

```go
type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
	Make()
}

type Tea struct{}
func (t *Tea) BoilWater()        { fmt.Println("Hirviendo agua para té") }
func (t *Tea) Brew()             { fmt.Println("Remojando la bolsita de té") }
func (t *Tea) PourInCup()        { fmt.Println("Sirviendo el té en la taza") }
func (t *Tea) AddCondiments()    { fmt.Println("Agregando limón") }
func (t *Tea) Make() {
	t.BoilWater()
	t.Brew()
	t.PourInCup()
	t.AddCondiments()
}
```

---

## 💻 Ejemplo en Go (Template Pattern)

```go
package main

import "fmt"

// Abstract Class
type DataParser interface {
	ReadData()
	ProcessData()
	SaveData()
	ParseTemplate()
}

// Template Method
func ParseTemplate(dp DataParser) {
	dp.ReadData()
	dp.ProcessData()
	dp.SaveData()
}

// Concrete Class: CSV
type CSVParser struct{}

func (c *CSVParser) ReadData()    { fmt.Println("Leyendo datos desde archivo CSV") }
func (c *CSVParser) ProcessData() { fmt.Println("Procesando datos del CSV") }
func (c *CSVParser) SaveData()    { fmt.Println("Guardando datos del CSV") }

func main() {
	csv := &CSVParser{}
	ParseTemplate(csv)
}

```

---

## 🛠️ Ejercicio Práctico

**Ejercicio:**

Implementa un sistema que modele el proceso de preparar distintas comidas (Pizza, Pasta, Ensalada).
Define los pasos genéricos como:

* Preparar ingredientes
* Cocinar (si aplica)
* Servir

Utiliza el patrón Template para organizar este flujo en una estructura base y personalizar los pasos específicos en cada comida.

**Objetivo:** reutilizar el flujo principal y permitir cambios flexibles en comidas distintas.

---

## 📚 Recursos adicionales

- [Refactoring Guru - Template Method](https://refactoring.guru/es/design-patterns/template-method)
- [DoFactory - Template Pattern](https://www.dofactory.com/net/template-method-design-pattern)
    
---

## 🚀 Desafío (opcional)

Integra el patrón Template Method en una simulación de una cafetería, con distintas bebidas (Té, Café, Chocolate) y condimentos opcionales.
Permite que el cliente elija qué bebida preparar y define un menú dinámico. Utiliza el patrón para controlar el flujo de preparación.

---

## 🔜 ¿Qué sigue?

En la siguiente clase veremos el patrón Visitor, que te permitirá separar algoritmos de las estructuras sobre las que operan, una técnica poderosa para extensibilidad sin modificar estructuras base.