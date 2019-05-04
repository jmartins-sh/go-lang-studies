package main

import "fmt"

type carro struct {
	name         string
	currentSpeed int
}

type ferrari struct {
	carro // anonymous field
	turbo bool
}

func main() {
	f := ferrari{}

	f.name = "F40"

	f.currentSpeed = 0

	f.turbo = true

	fmt.Printf("Is the Ferrari %s with It's turbo activated? %v\n", f.name, f.turbo)

	fmt.Println(f)
}
