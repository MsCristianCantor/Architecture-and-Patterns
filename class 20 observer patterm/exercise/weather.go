package main

import "fmt"

// Observer interface
type Observer interface {
	Update(temperature string)
}

type TemperatureDisplay struct {
	name string
}

func (c *TemperatureDisplay) Update(temperature string) {
	fmt.Printf("Hola TemperatureDisplay %s, la temperatura ahora es %s!\n", c.name, temperature)
}

type MobileApp struct {
	name string
}

func (c *MobileApp) Update(temperature string) {
	fmt.Printf("Hola MobileApp %s, la temperatura ahora es %s!\n", c.name, temperature)
}

// WeatherStationSubject interface
type WeatherStationSubject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify()
}

// Concrete WeatherStationSubject: WeatherStation
type WeatherStation struct {
	temperature string
	available   bool
	observers   []Observer
}

func NewWeatherStation(temperature string) *WeatherStation {
	return &WeatherStation{temperature: temperature}
}

func (c *WeatherStation) Register(observer Observer) {
	c.observers = append(c.observers, observer)
}

func (c *WeatherStation) Unregister(observer Observer) {
	for i, obs := range c.observers {
		if obs == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}

func (c *WeatherStation) Notify() {
	for _, observer := range c.observers {
		observer.Update(c.temperature)
	}
}

func (c *WeatherStation) ChangeTemperature(temperature string) {
	c.temperature = temperature
	fmt.Printf("\nOcurrio un cambio en la temperatura, la nueva temperatura es %s!\n", c.temperature)
	c.available = true
	c.Notify()
}

// Main
func main() {
	// Crear weather station
	product := NewWeatherStation("18")

	// Crear clientes
	customer1 := &TemperatureDisplay{name: "Carlos"}
	customer2 := &MobileApp{name: "Ana"}

	// Suscribir clientes al producto
	product.Register(customer1)
	product.Register(customer2)

	// El clima cambia
	product.ChangeTemperature("20")
	// El clima vuelva a cambiar
	product.ChangeTemperature("22")
}
