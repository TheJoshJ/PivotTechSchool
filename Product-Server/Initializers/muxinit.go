package Initializers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Connect struct {
	Router *mux.Router
}

func (c *Connect) MuxInit() {

	//create the router
	c.Router = mux.NewRouter()
	log.Println("Router Created")

	//load the handlers
	c.initializeRoutes()

	//load the server
	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", c.Router)
	if err != nil {
		log.Fatal(err)
	}

}
