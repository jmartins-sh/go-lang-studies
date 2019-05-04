package main

import "fmt"

func getValue(number int) int {

	defer fmt.Println("End!")

	if number >= 5000 {
		fmt.Println("High value...")

		return 5000
	}

	fmt.Println("Low value...")

	return number
}

func main() {
	getValue(200)
}
