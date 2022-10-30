package Initializers

import (
	"encoding/json"
	"log"
	"os"
	"product-server/models"
)

func ProductsInit() {
	var products []models.Product

	bs, err := os.ReadFile("objects/products.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bs, &products); err != nil {
		log.Fatal(err)
	}
}
