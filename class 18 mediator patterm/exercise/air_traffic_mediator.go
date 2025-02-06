package main

import "fmt"

// Mediator interface
type AirMediator interface {
	RequestTakeOFF(plane Plane)
	RequestLand(plane Plane)
	AddPlane(plane Plane)
}

// Concrete Mediator
type Airstrip struct {
	planes       []Plane
	usedAirstrip bool
}

func (c *Airstrip) RequestTakeOFF(plane Plane) {
	if c.usedAirstrip {
		plane.TakeOFF()
		c.usedAirstrip = false
	} else {
		plane.Wait()
	}
}

func (c *Airstrip) RequestLand(plane Plane) {
	if c.usedAirstrip {
		plane.Wait()
	} else {
		plane.Land()
		c.usedAirstrip = true
	}
}

func (c *Airstrip) AddPlane(plane Plane) {
	c.planes = append(c.planes, plane)
}

// Plane interface
type Plane interface {
	RequestTakeOFF()
	RequestLand()
	TakeOFF()
	Land()
	Wait()
}

// Concrete Colleague
type MediatorPlane struct {
	name     string
	mediator AirMediator
}

func (u *MediatorPlane) RequestTakeOFF() {
	fmt.Printf("%s Request Take OFF:", u.name)
	u.mediator.RequestTakeOFF(u)
}

func (u *MediatorPlane) RequestLand() {
	fmt.Printf("%s Request Land:", u.name)
	u.mediator.RequestLand(u)
}

func (u *MediatorPlane) TakeOFF() {
	fmt.Printf("%s Take OFF\n", u.name)
}

func (u *MediatorPlane) Land() {
	fmt.Printf("%s Land\n", u.name)
}

func (u *MediatorPlane) Wait() {
	fmt.Printf("%s Wait\n", u.name)
}

func main() {
	airstrip := &Airstrip{}

	plane1 := &MediatorPlane{name: "Plane1", mediator: airstrip}
	plane2 := &MediatorPlane{name: "Plane2", mediator: airstrip}
	plane3 := &MediatorPlane{name: "Plane3", mediator: airstrip}

	airstrip.AddPlane(plane1)
	airstrip.AddPlane(plane2)
	airstrip.AddPlane(plane3)

	plane1.RequestLand()
	plane2.RequestLand()
	plane3.RequestLand()
	plane1.RequestTakeOFF()
	plane2.RequestLand()
	plane3.RequestLand()
	plane2.RequestTakeOFF()
	plane3.RequestLand()
	plane1.RequestLand()
	plane3.RequestTakeOFF()
}
