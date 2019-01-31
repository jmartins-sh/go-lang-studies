package main

import "fmt"

func main() {
	// array is homogeneous struct and his length is fixed
	var grades [3]float64

	fmt.Println(grades)

	grades[0], grades[1], grades[2] = 7.8, 4.9, 9.8

	fmt.Println(grades)

	total := 0.0

	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}

	average := total / float64(len(grades))

	fmt.Printf("Average is %.2f\n", average)

	var names [5]string

	names[0],
		names[1],
		names[2],
		names[3],
		names[4] =
		"João Antonio",
		"Danielle De Castro",
		"Valentina Martins",
		"Aglaer Gertrudes Martins",
		"Dorvalina de Fatima Martins"

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
