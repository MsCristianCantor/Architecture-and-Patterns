package main

import "fmt"

// Clonable es la interfaz que define el método Clone
type Clonable interface {
	Clone() Clonable
}

// BaseConcretePrototype es una implementación concreta de Clonable
type BaseConcretePrototype struct {
	Name         string
	Value        int
	NumberString []string
}

type ConcretePrototypeWithSize struct {
	BaseConcretePrototype
	Size string
}

func (p *ConcretePrototypeWithSize) Clone() Clonable {
	clone := ConcretePrototypeWithSize{
		BaseConcretePrototype: BaseConcretePrototype{
			Name:  fmt.Sprintf("%s clone", p.Name),
			Value: p.Value,
		},
		Size: fmt.Sprintf("%s clone", p.Size),
	}
	for _, v := range p.NumberString {
		clone.NumberString = append(clone.NumberString, fmt.Sprintf("%s clone", v))
	}
	return &clone
}

type ConcretePrototypeWithDate struct {
	BaseConcretePrototype
	Date string
}

func (p *ConcretePrototypeWithDate) Clone() Clonable {
	clone := ConcretePrototypeWithDate{
		BaseConcretePrototype: BaseConcretePrototype{
			Name:  fmt.Sprintf("%s clone", p.Name),
			Value: p.Value,
		},
		Date: "2024-08-26",
	}
	for _, v := range p.NumberString {
		clone.NumberString = append(clone.NumberString, fmt.Sprintf("%s clone", v))
	}
	return &clone
}

func main() {
	// Crear un objeto prototypeWithSize
	prototypeWithSize := &ConcretePrototypeWithSize{
		BaseConcretePrototype: BaseConcretePrototype{
			Name:  "Prototype1",
			Value: 100,
			NumberString: []string{
				"1",
				"2",
				"3",
				"4",
			},
		},
		Size: "L",
	}

	// Clonar el prototypeWithSize
	cloneWithSize := prototypeWithSize.Clone().(*ConcretePrototypeWithSize)

	// Mostrar los resultados
	fmt.Printf(
		"Original Prototype with size: %s, %d, %v, %v\n",
		prototypeWithSize.Name,
		prototypeWithSize.Value,
		prototypeWithSize.NumberString,
		prototypeWithSize.Size,
	)
	fmt.Printf(
		"Cloned Prototype with size: %s, %d, %v, %v\n",
		cloneWithSize.Name,
		cloneWithSize.Value,
		cloneWithSize.NumberString,
		cloneWithSize.Size,
	)

	// Crear un objeto prototypeWithSize
	prototypeWithDate := &ConcretePrototypeWithDate{
		BaseConcretePrototype: BaseConcretePrototype{
			Name:  "Prototype1",
			Value: 100,
			NumberString: []string{
				"1",
				"2",
				"3",
				"4",
			},
		},
		Date: "N/a",
	}

	fmt.Println("----------")

	// Clonar el prototypeWithSize
	cloneWithDate := prototypeWithDate.Clone().(*ConcretePrototypeWithDate)

	// Mostrar los resultados
	fmt.Printf(
		"Original Prototype with date: %s, %d, %v, %v\n",
		prototypeWithDate.Name,
		prototypeWithDate.Value,
		prototypeWithDate.NumberString,
		prototypeWithDate.Date,
	)
	fmt.Printf(
		"Cloned Prototype with date: %s, %d, %v, %v\n",
		cloneWithDate.Name,
		cloneWithDate.Value,
		cloneWithDate.NumberString,
		cloneWithDate.Date,
	)
}
