package main

import (
	"fmt"
	"strings"
)

type person struct {
	name     string
	lastName string
}

func (p person) getFullName() string {
	return p.name + " " + p.lastName
}

func (p *person) setFullName(fullName string) {
	splittedName := strings.Split(fullName, " ")

	p.name = splittedName[0]

	p.lastName = splittedName[1]
}

func main() {
	newPerson := person{name: "João", lastName: "Martins"}

	fmt.Println(newPerson.getFullName())

	newPerson.setFullName("Luis Silva")

	fmt.Println(newPerson.getFullName())
}
