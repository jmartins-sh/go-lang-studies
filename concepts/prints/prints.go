package main

import "fmt"

func main() {

	fmt.Print("Same ")

	fmt.Print("line.")

	fmt.Println("Next text will be in a ")

	fmt.Println("New line")

	x := 3.149874

	// fmt.Println("The value of x is " + x) -> error mismatched type

	xs := fmt.Sprint(x) //to string

	fmt.Println("The value of x is " + xs) // works fine!

	fmt.Println("The value of x is", x) // works fine as well!

	fmt.Printf("The value of x is %.2f", x) // format interpolation

	a := 1
	b := 1.9999
	c := false
	d := "yep!"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	fmt.Printf("\n%v %v %v %v\n", a, b, c, d)

}
