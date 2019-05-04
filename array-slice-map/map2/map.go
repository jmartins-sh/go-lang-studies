package main

import "fmt"

func main() {

	emploeesSalary := map[string]float64{
		"João Antonio":     10520.65,
		"Gabriel Da Silva": 5000.30,
		"Jaqueline Santos": 2300,
	}

	emploeesSalary["Jucaquim"] = 200

	delete(emploeesSalary, "José")

	for name, salary := range emploeesSalary {
		fmt.Println(name, salary)
	}
}
