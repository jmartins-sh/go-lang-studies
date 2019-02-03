package main

import "fmt"

func main() {

	firstSlice := make([]int, 2, 5)

	secondSlice := append(firstSlice, 1, 2, 3)

	fmt.Println(firstSlice, secondSlice)

	firstSlice[0] = 5

	fmt.Println(firstSlice, secondSlice)
}
