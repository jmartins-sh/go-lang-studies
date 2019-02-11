package main

import "fmt"

func mult(a, b int) int {
	return a * b
}

func exec(function func(int, int) int, p1, p2 int) int {
	return function(p1, p2)
}

func main() {
	result := exec(mult, 2, 5)

	fmt.Println(result)
}
