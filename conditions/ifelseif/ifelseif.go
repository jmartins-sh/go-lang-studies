package main

import "fmt"

func valueToGrade(value float64) string {
	if value >= 9 && value <= 10 {
		return "A"
	} else if value >= 8 && value < 9 {
		return "B"
	} else if value >= 5 && value < 8 {
		return "C"
	} else if value >= 3 && value < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(valueToGrade(9.8))

	fmt.Println(valueToGrade(6.9))

	fmt.Println(valueToGrade(2.3))
}
