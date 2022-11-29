package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
)

type Database struct {
	DB *gorm.DB
}

type Product struct {
	ProductID   int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func main() {
	if err := checkForDB(); err != nil {
		log.Fatalf("Error checking for or removing DB\n%s", err)
	}
	db, err := newDB()
	if err != nil {
		log.Fatalf("unable to create databse\n%v", err)
	}

	defer db.DB.Close()

	db.DB.AutoMigrate(&Product{})

	fileByte, err := loadJSON()
	if err != nil {
		log.Fatalf("error loading JSON")
	}

	err = addToDB(fileByte, db.DB)
	if err != nil {
		log.Fatalf("unable to create databse\n%v", err)
	}
}

func checkForDB() error {
	if err := os.Remove("products.db"); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	log.Println("Existing DB Removed")
	return nil
}

func newDB() (Database, error) {
	db, err := gorm.Open("sqlite3", "products.db")
	if err != nil {
		return Database{}, err
	}
	log.Println("New DB successfully created")
	return Database{DB: db}, nil
}

func loadJSON() ([]byte, error) {
	jsonFile, err := os.Open("obj/products.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	file, err := os.ReadFile("obj/products.json")
	if err != nil {
		return nil, err
	}
	return file, nil
}

func addToDB(b []byte, db *gorm.DB) error {
	var products []Product

	if err := json.Unmarshal(b, &products); err != nil {
		return err
	}

	for _, product := range products {
		db.Create(&product)
	}

	for i := 1; i <= 5; i++ {
		db.Debug().Where(&Product{ProductID: i}).First(&products)
	}
	return nil
}
