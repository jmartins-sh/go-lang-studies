package main

import (
	"fmt"
	"reflect"
)

type dynamic interface{}

func main() {
	var stuff dynamic = "Hello World!"

	fmt.Println(reflect.TypeOf(stuff))

	stuff = 90.0

	fmt.Println(reflect.TypeOf(stuff))

	stuff = true

	fmt.Println(reflect.TypeOf(stuff))
}
