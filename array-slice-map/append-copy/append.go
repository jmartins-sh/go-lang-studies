package main

import "fmt"

func main() {

	theArray := [3]int{1, 2, 3}

	var theSlice []int

	fmt.Println(theArray, theSlice, cap(theSlice), len(theSlice))

	// Error -> theArray = append(theArray, 4, 5, 6)
	theSlice = append(theSlice, 4, 5, 6)

	theSecondSlice := make([]int, 2)

	copy(theSecondSlice, theSlice)

	fmt.Println(theSecondSlice)
}
