package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

func (p product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	var productOne product

	productOne = product{
		name:     "Pen",
		price:    1.80,
		discount: 0.05,
	}

	fmt.Println(productOne.priceWithDiscount())

	productTwo := product{"Notebook", 2000, 0.10}

	fmt.Println(productTwo.priceWithDiscount())
}
