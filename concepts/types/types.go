package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// integer numbers
	fmt.Println(1, 2, 1000)

	fmt.Println("Literal integer is", reflect.TypeOf(32000))

	// unsigned (only greater than zero 0) ... uint8 uint16 uint32 uint64
	var b byte = 255 // masquerade type (uint8)

	fmt.Println("O byte é", reflect.TypeOf(b))

	// signed (negatives and positives numbers) int8 int16 int32 int64
	integer := math.MaxInt64

	fmt.Println("The max value of a int is", integer)

	fmt.Println("The type of integer variable is", reflect.TypeOf(integer))

	//float numbers (float32, float64)
	var x float32 = 49.99

	fmt.Println("The type of x is", reflect.TypeOf(x))

	fmt.Println("The type of literal 49.99 is", reflect.TypeOf(49.99)) // float64 or float32 dependes on cpu architecture

	//booleans (true or false)
	bo := true

	fmt.Println("The type of bo is", reflect.TypeOf(bo))

	fmt.Println(!bo)

	//string
	s1 := "Hello my name is João"

	fmt.Println(s1 + "!")

	fmt.Println("The length of s1 is", len(s1))

	//string multiple lines
	s2 := `Hello
	my
	name
	is
	João`

	fmt.Println("The length of s1 is", len(s2))

	//char  ??
	char := 'a'

	fmt.Println("The type of char is", reflect.TypeOf(char)) // It's a int masquerade, ASCII

	fmt.Println(char)

	//rune ?
	// var run rune = 'b'

	// fmt.Println("Type of run is", reflect.TypeOf(run))

	// fmt.Println(run)
}
