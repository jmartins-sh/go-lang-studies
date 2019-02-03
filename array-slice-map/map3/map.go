package main

import "fmt"

func main() {
	emploeesGroupByFirstChar := map[string]map[string]float64{
		"J": {
			"João Antonio": 22000.36,
			"José Matheus": 15000,
		},
		"G": {
			"Gabriela Tavares": 10560,
			"Gustavo Almeida":  10000,
		},
		"P": {
			"Paulo Coelho": 500.36,
		},
	}

	delete(emploeesGroupByFirstChar, "P")

	for letter, emploees := range emploeesGroupByFirstChar {
		fmt.Println("The first letter in the name is", letter)

		fmt.Println("========================================")

		for name, salary := range emploees {
			fmt.Printf("Employe Name: %s (R$ %.2f)\n", name, salary)
		}

		fmt.Println("========================================")
	}
}
