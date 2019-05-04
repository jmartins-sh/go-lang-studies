package main

import "fmt"

func incValueByCopy(n int) {
	n++
}

func incValueByReference(n *int) {
	*n++
}

func main() {
	n := 1

	incValueByCopy(n)

	fmt.Println(n)

	incValueByReference(&n)

	fmt.Println(n)

}
