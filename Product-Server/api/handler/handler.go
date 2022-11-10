package api

import (
	"PivotTechSchool/Product-Server/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(&models.Products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("obj/products.json")
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

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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

	productsByte, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	file, err := os.Create("obj/products.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}(file)

	_, writeErr := file.Write(productsByte)
	if writeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {

	data, err := os.ReadFile("obj/products.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

func UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("obj/products.json")
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

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	price, _ := strconv.Atoi(r.FormValue("price"))
	if r.FormValue("price") == "" || r.FormValue("description") == "" || r.FormValue("name") == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newProduct := models.Product{
		ID:          id,
		Price:       price,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}

	for i, product := range products {
		if product.ID == id {
			products[i] = newProduct

			productsByte, err := json.Marshal(products)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			file, err := os.Create("obj/products.json")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}(file)

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

func DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("obj/products.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var products []models.Product
	err = json.Unmarshal(data, &products)

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)

			productsByte, err := json.Marshal(products)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			file, err := os.Create("obj/products.json")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}(file)

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
