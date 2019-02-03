package main

import "fmt"

func main() {

	s := make([]int, 10) // Creates a Slice with make() function, containing 10 positions

	s[9] = 12 // The tenth position has his value set to number twelve (12), we used the 9 to represent the Slice's index

	fmt.Println(s)

	s = make([]int, 2, 5) // Here we reuse the s Slice to reference a new Slice, which has his length to 2 elements and his capacity to 5

	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7) // Slice adapts dynamically when his max capacity is reached, this case the Slice doubled his capacity

	fmt.Println(s, len(s), cap(s))
}
