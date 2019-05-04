package main

import "fmt"

type sportCar interface {
	turnTurboOn()
}

type ferrari struct {
	model        string
	turboOn      bool
	currentSpeed int
}

func (f *ferrari) turnTurboOn() {
	f.turboOn = true
}

func main() {
	firstCar := ferrari{
		model:        "F40",
		turboOn:      false,
		currentSpeed: 0,
	}

	firstCar.turnTurboOn()

	var secondCar sportCar = &ferrari{
		model:        "F40",
		turboOn:      false,
		currentSpeed: 0,
	}

	secondCar.turnTurboOn()

	fmt.Println(firstCar, secondCar)
}
