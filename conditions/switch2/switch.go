package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch { // switch looking for the first case with expression which returns true
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
