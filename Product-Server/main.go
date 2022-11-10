package main

import (
	"PivotTechSchool/Product-Server/api/handler"
	"PivotTechSchool/Product-Server/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type server struct {
	Router *mux.Router
}

var Products []models.Product

func main() {
	s := server{}
	s.Router = mux.NewRouter()

	initProducts("obj/products.json")

	log.Println("Router Created")
	log.Println("Loading Routes...")

	s.Router.HandleFunc("/products", api.GetProducts).Methods("GET")
	s.Router.HandleFunc("/products", api.AddProduct).Methods("POST")
	s.Router.HandleFunc("/products/{id}", api.GetProductByID).Methods("GET")
	s.Router.HandleFunc("/products/{id}", api.UpdateProductByID).Methods("PUT")
	s.Router.HandleFunc("/products/{id}", api.DeleteProductByID).Methods("DELETE")
	log.Println("Loaded Routes!")

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}

func initProducts(filepath string) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to read products list - %s.", err)
		return
	}
	if err = json.Unmarshal(data, &Products); err != nil {
		log.Fatalf("Error unmarshalling products list - %s", err)
	}
}
