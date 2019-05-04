package main

import "fmt"

func valueToGrade(value float64) string {
	val := int(value)

	switch val {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid Grade!"
	}
}

func main() {
	fmt.Println(valueToGrade(9.5))
	fmt.Println(valueToGrade(6.2))
	fmt.Println(valueToGrade(3.8))
	fmt.Println(valueToGrade(11))
}
