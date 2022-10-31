package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"product-server/api/handler"
)

type Connect struct {
	Router *mux.Router
}

func main() {
	c := Connect{}
	c.Router = mux.NewRouter()
	log.Println("Router Created")

	log.Println("Loading Routes...")

	c.Router.HandleFunc("/products", api.GetProducts).Methods("GET")
	c.Router.HandleFunc("/products", api.AddProduct).Methods("POST")
	c.Router.HandleFunc("/products/{id}", api.GetProductByID).Methods("GET")
	c.Router.HandleFunc("/products/{id}", api.UpdateProductByID).Methods("PUT")
	c.Router.HandleFunc("/products/{id}", api.DeleteProductByID).Methods("DELETE")
	log.Println("Loaded Routes!")

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", c.Router)
	if err != nil {
		log.Fatal(err)
	}
}
