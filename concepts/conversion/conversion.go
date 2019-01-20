package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4

	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9

	notaFinal := int(nota)

	fmt.Println(notaFinal)

	// careful
	fmt.Println("Teste " + string(97))

	// int to string
	fmt.Println("Teste " + strconv.Itoa(97))

	// string to int
	num, _ := strconv.Atoi("97")

	fmt.Println(num - 96)

	// to bool
	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("Verdadeiro")
	}
}
