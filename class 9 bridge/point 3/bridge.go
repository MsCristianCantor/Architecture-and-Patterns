package main

import "fmt"

type AbstractionDevice interface {
	PowerOn()
	PowerOff()
}

type TvDevice struct {
}

func (a *TvDevice) PowerOn() {
	fmt.Println("TV Power On")
}

func (a *TvDevice) PowerOff() {
	fmt.Println("TV Power Off")
}

type RadioDevice struct {
}

func (a *RadioDevice) PowerOn() {
	fmt.Println("Radio Power On")
}

func (a *RadioDevice) PowerOff() {
	fmt.Println("Radio Power Off")
}

type ImplementationRemoteControl interface {
	SetDevice(AbstractionDevice)
	TurnOn()
	TurnOff()
}

type BasicRemoteControl struct {
	Device AbstractionDevice
}

func (i *BasicRemoteControl) SetDevice(device AbstractionDevice) {
	i.Device = device
}

func (i *BasicRemoteControl) TurnOn() {
	fmt.Println("Turn On via basic remote")
	i.Device.PowerOn()
}

func (i *BasicRemoteControl) TurnOff() {
	fmt.Println("Turn Off via basic remote")
	i.Device.PowerOff()
}

type AdvancedRemoteControl struct {
	Device AbstractionDevice
}

func (i *AdvancedRemoteControl) SetDevice(device AbstractionDevice) {
	i.Device = device
}

func (i *AdvancedRemoteControl) TurnOn() {
	fmt.Println("Turn On via advanced remote")
	i.Device.PowerOn()
}

func (i *AdvancedRemoteControl) TurnOff() {
	fmt.Println("Turn Off via advanced remote")
	i.Device.PowerOff()
}

func main() {
	var tvDevice AbstractionDevice = &TvDevice{}
	var radioDevice AbstractionDevice = &RadioDevice{}

	var basicRemoteControl ImplementationRemoteControl = &BasicRemoteControl{
		Device: tvDevice,
	}
	var advancedRemoteControl ImplementationRemoteControl = &AdvancedRemoteControl{
		Device: tvDevice,
	}

	// Basic device with tv
	basicRemoteControl.TurnOn()
	basicRemoteControl.TurnOff()

	// Advanced device with tv
	advancedRemoteControl.TurnOn()
	advancedRemoteControl.TurnOff()

	// overwrite device
	basicRemoteControl.SetDevice(radioDevice)
	advancedRemoteControl.SetDevice(radioDevice)

	// Basic device with radio
	basicRemoteControl.TurnOn()
	basicRemoteControl.TurnOff()

	// Advanced device with radio
	advancedRemoteControl.TurnOn()
	advancedRemoteControl.TurnOff()
}
