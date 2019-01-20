package main

import "fmt"

func main() {
	var a int     // 0
	var b float64 // 0
	var c bool    // false
	var d string  // ""	(empty string)
	var e *int    // nil (reference type)

	fmt.Printf("%v %v %v %v %v", a, b, c, d, e)
}
