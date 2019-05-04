package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name     string
	lastName string
}

func (p person) toString() string {
	return p.name + " " + p.lastName
}

type product struct {
	name  string
	price float64
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {
	var newPerson printable = person{name: "João Antonio", lastName: "Martins Filho"}

	var newProduct printable = product{name: "Notebook", price: 7250.60}

	fmt.Println(newPerson.toString())

	print(newPerson)

	fmt.Println(newProduct.toString())

	print(newProduct)

}
