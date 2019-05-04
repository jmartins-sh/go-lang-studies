package main

import "fmt"

func main() {
	x := 1
	y := 2
	z := true

	//only postfix operators
	x++ // increment
	fmt.Println(x)

	y-- // decrement
	fmt.Println(y)

	fmt.Println(!z) // negation
}
