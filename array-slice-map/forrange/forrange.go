package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5} //compiler counts the indexes [...] represent a array

	// "foreach loop" in go lang
	for i, number := range numbers {
		fmt.Printf("%d) %d\n", i, number)
	}

	// underscore just ignore the variable
	for _, number := range numbers {
		fmt.Println(number)
	}
}
