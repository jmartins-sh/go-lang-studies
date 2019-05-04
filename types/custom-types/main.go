package main

import "fmt"

type grade float64

func (g grade) between(begin, end float64) bool {
	return begin <= float64(g) && end >= float64(g)
}

func gradeToLetter(g grade) (letter string) {

	switch {
	case g.between(9.0, 10.0):
		letter = "A"
		break
	case g.between(7.0, 8.99):
		letter = "B"
		break
	case g.between(5.0, 6.99):
		letter = "C"
		break
	case g.between(3.0, 4.99):
		letter = "D"
		break
	case g.between(0.0, 2.99):
		letter = "E"
		break
	default:
		letter = "Invalida Grade, verify!"
	}

	return
}

func main() {
	fmt.Println(gradeToLetter(9.8))
	fmt.Println(gradeToLetter(6.9))
	fmt.Println(gradeToLetter(2.1))
	fmt.Println(gradeToLetter(-1))
}
