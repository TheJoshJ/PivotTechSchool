package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Connect struct {
	DB     *gorm.DB
	log    *log.Logger
	Router *mux.Router
}

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func main() {
	var err error
	c := Connect{}
	c.Router = mux.NewRouter()

	c.DB, err = initConnection("../db-products-server/seeder/products.db")
	if err != nil {
		log.Printf("Unable to connect to databse source\n%v", err)
		os.Exit(1)
	}

	log.Println("Router Created")
	log.Println("Loading Routes...")

	c.Router.HandleFunc("/products", c.GetProducts).Methods("GET")
	c.Router.HandleFunc("/products", c.AddProduct).Methods("POST")
	c.Router.HandleFunc("/products/{id}", c.GetProductByID).Methods("GET")
	c.Router.HandleFunc("/products/{id}", c.UpdateProductByID).Methods("PUT")
	c.Router.HandleFunc("/products/{id}", c.DeleteProductByID).Methods("DELETE")
	log.Println("Loaded Routes!")

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", c.Router))
}

func initConnection(filepath string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	log.Println("New DB successfully created")
	return db, nil
}

func (c *Connect) GetProducts(w http.ResponseWriter, r *http.Request) {

	limit := r.URL.Query().Get("limit")
	if limit == "" {
		// id.asc is the default sort query
		limit = "20"
	}

	sortBy := r.URL.Query().Get("sort")
	if sortBy == "" {
		// id.asc is the default sort query
		sortBy = "id"
	}

	limitInt, _ := strconv.Atoi(limit)
	if limitInt > 100 {
		limitInt = 100
	}
	if limitInt < 0 {
		limitInt = 20
	}
	switch sortBy {
	case "id":
	case "name":
	case "price":
	default:
		sortBy = "id"
	}

	var productsRet []Product
	c.DB.Table("products").Select("*").Limit(limitInt).Order(sortBy).Scan(&productsRet)

	reply, err := json.Marshal(&productsRet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(reply)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Connect) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if product.ID == 0 || product.Name == "" || product.Price == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var productRet Product
	c.DB.Table("products").First(&productRet, "ID = ?", product.ID).Scan(&productRet)
	if productRet.ID == product.ID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c.DB.Create(&product)
	w.WriteHeader(http.StatusCreated)
}

func (c *Connect) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, err := strconv.Atoi(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var productRet Product
	c.DB.Table("products").First(&productRet, "ID = ?", id).Scan(&productRet)

	if productRet.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	reply, err := json.Marshal(&productRet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(reply)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Connect) UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, err := strconv.Atoi(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if product.ID == 0 || product.Name == "" || product.Price == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//test to make sure that the product exists
	var productRet Product
	c.DB.Table("products").First(&productRet, "ID = ?", id).Scan(&productRet)
	if productRet.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//if the product ID in the URL and the req don't match, don't accept the change
	idInt, err := strconv.Atoi(id)
	if idInt != product.ID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//update the product
	c.DB.Model(&product).Where("id = ?", idInt).Updates(Product{Name: product.Name, Description: product.Description, Price: product.Price, ID: product.ID})
	w.WriteHeader(http.StatusOK)
}

func (c *Connect) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, err := strconv.Atoi(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if idInt == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//check to make sure the product exists
	var productCheck Product
	c.DB.Table("products").First(&productCheck, "ID = ?", id).Scan(&productCheck)
	if productCheck.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//delete the item if it exists
	c.DB.Where("ID = ?", id).Delete(&Product{})

	//check to ensure that the item was deleted
	var productRet Product
	c.DB.Table("products").First(&productRet, "ID = ?", id).Scan(&productRet)
	if productRet.ID == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}
