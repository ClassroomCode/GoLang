package main

import (
	"log"
	"net/http"

	"example.com/service/product/entities"
	"example.com/service/product/httphandlers"
	"example.com/service/product/storage"
)

var productID = 0

func createProduct(name string, price float64) entities.Product {
	productID++
	return entities.Product{
		ID:    productID,
		Name:  name,
		Price: price,
	}
}

func main() {
	log.Println("Creating dummy products")

	storage.Add(createProduct("Testing", 99.95))
	storage.Add(createProduct("Testing Again", 149.95))
	storage.Add(createProduct("Testing A Third Time", 29.00))

	log.Println("Attempting to start HTTP Server.")

	http.HandleFunc("/", httphandlers.HandleRequest)

	var err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Panicln("Server failed starting. Error: %w", err)
	}
}
