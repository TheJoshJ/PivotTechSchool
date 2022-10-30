package Initializers

import (
	"log"
	"product-server/Handlers"
)

func (c *Connect) initializeRoutes() {
	c.Router.HandleFunc("/products", Handlers.GetProducts).Methods("GET")
	c.Router.HandleFunc("/products", Handlers.PostProducts).Methods("POST")
	c.Router.HandleFunc("/products/{id}", Handlers.GetProductByID).Methods("GET")
	c.Router.HandleFunc("/products/{id}", Handlers.PutProductByID).Methods("PUT")
	c.Router.HandleFunc("/products/{id}", Handlers.DeleteProductByID).Methods("DELETE")

	log.Println("Loaded Routes")
}
