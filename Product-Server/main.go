package main

import (
	"product-server/Initializers"
)

func main() {
	c := Initializers.Connect{}
	c.MuxInit()
	Initializers.ProductsInit()
}
