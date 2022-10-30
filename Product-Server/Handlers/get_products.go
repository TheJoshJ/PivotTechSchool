package Handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"product-server/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("objects/products.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var products []models.Product
	err = json.Unmarshal(data, &products)

	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
