package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1

	for i <= 10 { // only loop in Go
		fmt.Println(i)

		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d\n", i)
	}

	fmt.Println("Mixing!")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

	for {
		fmt.Println("Infinite loop...")

		time.Sleep(time.Second * 5)
	}

}
