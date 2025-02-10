package main

import "fmt"

// Memento almacena el estado de un objeto
 type Memento struct {
    state string
}

// Originator crea y usa Mementos
 type Originator struct {
    state string
}

func (o *Originator) Save() *Memento {
    return &Memento{state: o.state}
}

func (o *Originator) Restore(m *Memento) {
    o.state = m.state
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
    originator := &Originator{state: "Estado Inicial"}
    caretaker := &Caretaker{}

    // Guardar estado
    caretaker.AddMemento(originator.Save())
    originator.state = "Estado Modificado"
    fmt.Println("Estado actual:", originator.state)

    // Restaurar estado anterior
    originator.Restore(caretaker.GetMemento(0))
    fmt.Println("Estado restaurado:", originator.state)
}