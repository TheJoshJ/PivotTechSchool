package Handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"product-server/models"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {

	data, err := os.ReadFile("objects/products.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var products []models.Product
	err = json.Unmarshal(data, &products)

	for _, product := range products {
		if product.ID == id {
			err := json.NewEncoder(w).Encode(product)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	return
}
