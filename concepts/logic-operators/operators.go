package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	buyTv50 := trab1 && trab2 // AND

	buyTv32 := trab1 != trab2 // XOR (sample)

	buyIceCream := trab1 || trab2 // OR

	return buyTv50, buyTv32, buyIceCream
}

func main() {
	tv50, tv32, iceCream := compras(true, true)

	fmt.Printf("Tv50: %t, Tv32: %t, IceCream: %t, Healthy: %t\n\n",
		tv50, tv32, iceCream, !iceCream)

	tv50, tv32, iceCream = compras(false, true)

	fmt.Printf("Tv50: %t, Tv32: %t, IceCream: %t, Healthy: %t\n\n",
		tv50, tv32, iceCream, !iceCream)

	tv50, tv32, iceCream = compras(false, false)

	fmt.Printf("Tv50: %t, Tv32: %t, IceCream: %t, Healthy: %t\n\n",
		tv50, tv32, iceCream, !iceCream)
}
