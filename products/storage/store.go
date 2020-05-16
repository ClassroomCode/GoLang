package storage

import (
	"example.com/service/product/entities"
)

var store entities.ProductList
var currentMaxID = 1

// Get retrieves all of the products
func Get() entities.ProductList {
	return store
}

// Add adds a product and returns the ID
func Add(product entities.Product) int {
	product.ID = currentMaxID
	currentMaxID++
	store = append(store, product)
	return product.ID
}

// Remove removes a product from the store
func Remove(id int) bool {
	index := -1

	for i, product := range store {
		if product.ID == id {
			index = i
		}
	}

	if index != -1 {
		store = append(store[:index], store[index+1:]...)
	}

	// Returns true if item was found & removed
	return index != -1
}
