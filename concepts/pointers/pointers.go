package main

import "fmt"

func main() {
	i := 1

	var p *int // creating a point with "*" before the type

	p = &i // getting the address of i variable

	// p++ Not allowed

	*p++ // incrementing the value from reference in pointer

	i++ // incrementing the value from variable

	//*p and i has the same value, because *p is the pointer to i variable value
	// p has the reference of i, so p is equal to &i (& operator in this case return the reference of i (variable))
	fmt.Printf("p: %v\n&i: %v\n*p: %d\ni: %d", p, &i, *p, i)
}
