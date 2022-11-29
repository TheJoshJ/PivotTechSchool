package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
)

type Database struct {
	DB  *gorm.DB
	log *log.Logger
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

	err = addToDB(fileByte, db)
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

func addToDB(b []byte, db Database) error {
	var products []Product
	var productsReturn []Product

	if err := json.Unmarshal(b, &products); err != nil {
		return err
	}

	for _, product := range products {
		db.DB.Create(&product)
	}

	db.DB.Table("products").Select("*").Limit(5).Scan(&productsReturn)
	log.Println(productsReturn)

	/*
		OUTPUT:
		[{1 Water - San Pellegrino curae nulla dapibus dolor vel est donec odio justo sollicitudin ut suscipit a feugiat et eros vestibulum ac est lacinia 80} {2 Cape Capensis - Fillet rutrum neque aenean auctor gravida
		sem praesent id massa id nisl venenatis lacinia aenean sit amet justo morbi ut odio 53} {3 Bread - Bistro Sour ac est lacinia nisi venenatis tristique fusce congue diam id 22} {5 Lettuce - Arugula vel nisl duis ac nibh fusce lacus p
		urus aliquet at feugiat non pretium 11} {6 Kahlua porttitor lorem id ligula suspendisse ornare consequat lectus in est risus auctor sed tristique in tempus sit 82}]
	*/

	return nil
}
