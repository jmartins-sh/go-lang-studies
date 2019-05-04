package main

import "fmt"

func main() {
	var b byte = 3

	fmt.Println(b)

	i := 3 // type inference
	i += 3 // add -> i = i + 3
	i -= 3 // sub -> i = i - 3
	i /= 2 // div -> i = i / 3
	i *= 2 // mul -> i = i * 3
	i %= 2 // mod -> i = i % 3

	fmt.Println(i)

	x, y := 1, 2

	fmt.Println(x, y)

	x, y = y, x // easy way to switch values

	fmt.Println(x, y)
}
