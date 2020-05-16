package httphandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"log"

	"example.com/service/product/entities"
	"example.com/service/product/httphandlers/httputils"
	"example.com/service/product/storage"
)

// Add uses incoming JSON request to add a new product
func Add(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	byteData, err := ioutil.ReadAll(r.Body)

	if err != nil {
		httputils.HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
		return
	}

	var product entities.Product

	err = json.Unmarshal(byteData, &product)

	if err != nil {
		httputils.HandleError(&w, 500, "Internal Server Error", "Error unmarhsalling JSON", err)
		return
	}

	if product.Name == "" || product.Price < 1 {
		httputils.HandleError(&w, 400, "Bad Request", "Unmarshalled JSON didn't have required fields", nil)
		return
	}

	id := storage.Add(product)

	log.Println("Added product:", product)

	httputils.HandleSuccess(&w, entities.ID{ID: id})
}
