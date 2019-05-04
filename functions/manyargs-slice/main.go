package main

import "fmt"

func printApproved(candidates ...string) {
	fmt.Println("Approved List")

	for i, approved := range candidates {
		fmt.Printf("%d) %s\n", i+1, approved)
	}
}

func main() {
	approveds := []string{"Maria", "Pedro", "Rafael", "João", "Guilherme"}

	printApproved(approveds...)
}
