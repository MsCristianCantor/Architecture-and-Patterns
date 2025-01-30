package main

import "fmt"

// Command Interface
type Command interface {
	Execute()
	Undo()
}

// Receiver: Light
type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("Luz encendida")
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("Luz apagada")
}

// ConcreteCommand: LightOnCommand
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

func (c *LightOnCommand) Undo() {
	c.light.TurnOff()
}

// ConcreteCommand: LightOffCommand
type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}

func (c *LightOffCommand) Undo() {
	c.light.TurnOn()
}

// Invoker: RemoteControl
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func (r *RemoteControl) PressUndo() {
	r.command.Undo()
}

func main() {
	// Receiver
	light := &Light{}

	// Commands
	lightOn := &LightOnCommand{light: light}
	lightOff := &LightOffCommand{light: light}

	// Invoker
	remote := &RemoteControl{}

	// Encender la luz
	remote.SetCommand(lightOn)
	remote.PressButton()
	remote.PressUndo()

	// Apagar la luz
	remote.SetCommand(lightOff)
	remote.PressButton()
	remote.PressUndo()
}
