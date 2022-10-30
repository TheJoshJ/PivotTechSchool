package Handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"product-server/models"
	"strconv"
)

func PostProducts(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("objects/products.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var products []models.Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	if r.FormValue("price") == "" || r.FormValue("description") == "" || r.FormValue("name") == "" || r.FormValue("id") == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, product := range products {
		if product.ID == id {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	newProduct := models.Product{
		ID:          id,
		Price:       price,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}

	products = append(products, newProduct)

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
}
