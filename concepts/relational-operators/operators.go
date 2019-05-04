package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("==", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Dates: ", d1 == d2) // True, structs are value type
	fmt.Println("Dates with Equal: ", d1.Equal(d2))

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}

	fmt.Println("Pessoas: ", p1 == p2)
}

type Pessoa struct {
	Nome string
}
