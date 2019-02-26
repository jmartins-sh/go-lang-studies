package main

import "fmt"

func init() { // Runs first before main function
	fmt.Println("Init...")
}

func main() {
	fmt.Println("Main...")
}
