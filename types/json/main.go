package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"product_id"`
	Name  string   `json:"product_name"`
	Price float64  `json:"product_price"`
	Tags  []string `jsong:"product_tags"`
}

func main() {
	// struct to json
	newProduct := product{ID: 1, Name: "Notebook", Price: 1250.62, Tags: []string{"IT", "Internet"}}

	newProductJSON, _ := json.Marshal(newProduct)

	fmt.Println(string(newProductJSON))

	//json to struct
	var secondProduct product

	jsonString := `{"product_id":2, "product_name": "Pen", "product_price": 1.58, "product_tags":["Office","School"]}`

	json.Unmarshal([]byte(jsonString), &secondProduct)

	fmt.Println(secondProduct)
}
