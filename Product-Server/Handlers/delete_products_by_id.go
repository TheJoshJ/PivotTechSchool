package Handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"product-server/models"
	"strconv"
)

func DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("objects/products.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var products []models.Product
	err = json.Unmarshal(data, &products)

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)

			productsByte, _ := json.MarshalIndent(products, "", " ")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			file, err := os.Create("objects/products.json")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			defer file.Close()

			_, writeErr := file.Write(productsByte)
			if writeErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}
