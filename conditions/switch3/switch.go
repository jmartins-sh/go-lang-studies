package main

import "fmt"

func valueToGrade(value float64) string {
	switch {
	case value >= 9 && value <= 10:
		return "A"
	case value >= 8 && value < 9:
		return "B"
	case value >= 5 && value < 8:
		return "C"
	case value >= 3 && value < 5:
		return "D"
	case value >= 0 && value < 3:
		return "E"
	default:
		return "Invalid Grade!"
	}
}

func main() {
	fmt.Println(valueToGrade(9.8))
	fmt.Println(valueToGrade(7.4))
	fmt.Println(valueToGrade(5.6))
	fmt.Println(valueToGrade(4.2))
	fmt.Println(valueToGrade(2.9))
	fmt.Println(valueToGrade(11))
}
