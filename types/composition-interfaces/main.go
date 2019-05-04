package main

import "fmt"

type sportCar interface {
	turnTurboOn()
}

type lustCar interface {
	park()
}

type sportLustCar interface {
	sportCar
	lustCar
	// if you wanna add some more functions, be my guest
}

type bmw struct {
}

func (b bmw) turnTurboOn() {
	fmt.Println("Turning turbo on...")
}

func (b bmw) park() {
	fmt.Println("Parking...")
}

func main() {
	var b sportLustCar = bmw{}

	b.turnTurboOn()

	b.park()
}
