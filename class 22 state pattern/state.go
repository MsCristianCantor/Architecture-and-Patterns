package main

import "fmt"

// Estado define la interfaz que todos los estados deben implementar
type State interface {
	Handle(context *Context)
}

// Context mantiene una referencia al estado actual
type Context struct {
	state State
}

func (c *Context) SetState(s State) {
	c.state = s
}

func (c *Context) Request() {
	c.state.Handle(c)
}

// Estado concreto A
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle(c *Context) {
	fmt.Println("Estado A: cambiando al estado B")
	c.SetState(&ConcreteStateB{})
}

// Estado concreto B
type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle(c *Context) {
	fmt.Println("Estado B: cambiando al estado A")
	c.SetState(&ConcreteStateA{})
}

// Main
func main() {
	context := &Context{}
	stateA := &ConcreteStateA{}
	context.SetState(stateA)

	// Llamadas sucesivas cambian el estado
	context.Request()
	context.Request()
	context.Request()
}
