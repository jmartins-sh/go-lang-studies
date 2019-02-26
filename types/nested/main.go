package main

import "fmt"

type item struct {
	ID    int
	qtd   int
	preco float64
}

type order struct {
	ID     int
	userID int
	Itens  []item
}

func (o order) total() (total float64) {

	for _, item := range o.Itens {
		total += item.preco * float64(item.qtd)
	}

	return
}

func main() {
	newOrder := order{
		ID:     1,
		userID: 1,
		Itens: []item{
			item{ID: 1, preco: 12.50, qtd: 3},
			item{ID: 2, preco: 27.50, qtd: 4},
			item{ID: 3, preco: 1.50, qtd: 8},
		},
	}

	fmt.Printf("Total Order Price: R$ %.2f", newOrder.total())
}
