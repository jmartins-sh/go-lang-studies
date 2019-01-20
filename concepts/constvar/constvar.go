package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415 // constant with initial value assign and type defined

	var raio = 3.2 // type inference by assignment

	area := PI * math.Pow(raio, 2) // reduced form to declare and assign a value to a variable

	fmt.Println("The circle's area is", area) // printing in CLI (Console) with fmt package

	const ( // declaration block - constant
		A = 1
		B = 2
	)

	var ( // declaration block - variables
		c = 3
		d = 4
	)

	// remember * Go doesn't allow declare a variable and not use It
	// you have to use in somewhere in the code, otherwise will result a error

	fmt.Println(A, B, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "Ya!"

	fmt.Println(g, h, i)
}
