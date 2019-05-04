package main

import "fmt"

// There's no ternary operator in Go Lang
func getResult(studentNote float64) string {
	// the same to return studentNote >= 7 ? "Approved" : "Disapproved"
	if studentNote >= 7 {
		return "Approved"
	}

	return "Disapproved"
}

func main() {
	fmt.Println(getResult(7))
}
