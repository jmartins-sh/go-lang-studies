package main

import "fmt"

func main() {

	// var ordersByInvoice map[string]int
	// Maps must be initialize before assign any value to them
	ordersByInvoice := make(map[string]int)

	ordersByInvoice["42589-1"] = 20001

	ordersByInvoice["43789-0"] = 20227

	ordersByInvoice["42987-3"] = 20298

	fmt.Println(ordersByInvoice) // Printing all the elements in map

	for invoice, order := range ordersByInvoice { // using loop to go through the map
		fmt.Printf("NF: %s - Pedido: %d\n", invoice, order)
	}

	fmt.Println(ordersByInvoice["42987-3"])

	delete(ordersByInvoice, "42987-3")

	fmt.Println(ordersByInvoice["42987-3"])
}
