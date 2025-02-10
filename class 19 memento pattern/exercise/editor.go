package main

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
