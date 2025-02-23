package main

import "fmt"

// Memento almacena el estado de un objeto
type Memento struct {
	texto string
}

// Editor crea y usa Mementos
type Editor struct {
	texto string
}

func (o *Editor) Save() *Memento {
	return &Memento{texto: o.texto}
}

func (o *Editor) Restore(m *Memento) {
	o.texto = m.texto
}

// Caretaker administra los Mementos
type Caretaker struct {
	history []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.history = append(c.history, m)
}

func (c *Caretaker) GetMemento(index int) *Memento {
	if index < len(c.history) {
		return c.history[index]
	}
	return nil
}

func main() {
	editor := &Editor{texto: "Texto inicial"}
	caretaker := &Caretaker{}
	fmt.Println("Texto inicial:", editor.texto)

	// Guardar estado
	caretaker.AddMemento(editor.Save())
	editor.texto = "Texto modificado"
	fmt.Println("Texto actual:", editor.texto)

	// Restaurar estado anterior
	editor.Restore(caretaker.GetMemento(0))
	fmt.Println("Texto restaurado:", editor.texto)
}
