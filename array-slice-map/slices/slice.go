package main

import (
	"fmt"
	"reflect"
)

func main() {

	a1 := [3]int{1, 2, 3} // array

	s1 := []int{1, 2, 3} // slice

	fmt.Println(a1, s1)

	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//Slice isn't a array! Slide defines a part from one specific array
	s2 := a2[1:3]

	fmt.Println(a2, s2)

	s3 := a2[:2] // New Slice, but It points to the same array as the above

	fmt.Println(s3)

	// You can imagine a Slice like a struct having a length and a pointer which points to a element inside another array
	s4 := s2[:1]

	fmt.Println(s4)
}
