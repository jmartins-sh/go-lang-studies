package main

import "fmt"

func switcher(param1, param2 int) (first, second int) {
	first = param2

	second = param1

	return
}

func main() {
	x, y := switcher(9, 4)

	fmt.Println(x, y)
}
