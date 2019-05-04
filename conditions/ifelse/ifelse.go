package main

import "fmt"

func printResult(studentGrade float64) {
	if studentGrade >= 7 {
		fmt.Println("Approved with grade", studentGrade)
	} else {
		fmt.Println("Disapproved with grade", studentGrade)
	}
}

func main() {
	printResult(7.3)

	printResult(5.9)
}
