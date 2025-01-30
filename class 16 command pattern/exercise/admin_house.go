package main

import "fmt"

// Command Interface
type Command interface {
	Execute()
	Undo()
}

// Receiver: Music
type Music struct {
	status string
}

func (m *Music) Play() {
	m.status = "play"
	fmt.Println("Music played")
}

func (m *Music) Pause() {
	m.status = "pause"
	fmt.Println("Music paused")
}

func (m *Music) Stop() {
	m.status = "stop"
	fmt.Println("Music stoped")
}

// Receiver: Fan
type Fan struct {
	isOn  bool
	speed int
}

func (f *Fan) TurnON() {
	f.isOn = true
	fmt.Println("Turn on fan")
}

func (f *Fan) TurnOFF() {
	f.isOn = false
	fmt.Println("Turn off fan")
	f.ChangeSpeed(0)
}

func (l *Fan) ChangeSpeed(speed int) {
	l.speed = speed
	fmt.Println("Speed Changed, current speed: ", l.speed)
}

// ConcreteCommand: TurnOnCommand
type TurnOnCommand struct {
	music *Music
	fan   *Fan
}

func (c *TurnOnCommand) Execute() {
	c.music.Play()
	c.fan.TurnON()
	c.fan.ChangeSpeed(50)
}

func (c *TurnOnCommand) Undo() {
	c.music.Stop()
	c.fan.TurnOFF()
}

// ConcreteCommand: TurnOffCommand
type TurnOffCommand struct {
	music *Music
	fan   *Fan
}

func (c *TurnOffCommand) Execute() {
	c.music.Stop()
	c.fan.TurnOFF()
}

func (c *TurnOffCommand) Undo() {
	c.music.Play()
	c.fan.TurnON()
	c.fan.ChangeSpeed(50)
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
	music := &Music{}
	fan := &Fan{}

	// Commands
	turnOn := &TurnOnCommand{music: music, fan: fan}
	turnOff := &TurnOffCommand{music: music, fan: fan}

	// Invoker
	remote := &RemoteControl{}

	// Encender la Casa
	remote.SetCommand(turnOn)
	remote.PressButton()
	remote.PressUndo()

	fmt.Println("----------------")

	// Apagar la Casa
	remote.SetCommand(turnOff)
	remote.PressButton()
	remote.PressUndo()
}
