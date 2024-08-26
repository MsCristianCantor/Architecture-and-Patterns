package main

import (
	"fmt"
	"sync"
)

const default_time = 100

// Singleton estructura que queremos asegurar que solo tenga una instancia
type Singleton struct {
	name string
	time int
}

var instance *Singleton
var once sync.Once

// GetInstance retorna la instancia única del Singleton
func GetInstance(time int) *Singleton {
	once.Do(func() {
		instance = &Singleton{
			name: "UniqueInstance",
			time: default_time,
		}
	})
	instance.time = time
	return instance
}

func main() {
	singleton1 := GetInstance(200)
	fmt.Println(singleton1)
	singleton2 := GetInstance(300)
	fmt.Println(singleton2)

	if singleton1 == singleton2 {
		fmt.Println("Both variables point to the same instance.")
	}
}
