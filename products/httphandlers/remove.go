package httphandlers

import (
	"io/ioutil"
	"net/http"

	"encoding/json"

	"example.com/service/product/entities"
	"example.com/service/product/httphandlers/httputils"
	"example.com/service/product/storage"
)

// Remove removes a product from the store
func Remove(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		httputils.HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
		return
	}

	var id entities.ID

	err = json.Unmarshal(requestBody, &id)

	if err != nil {
		httputils.HandleError(&w, 400, "Bad Request", "Error unmarshalling", err)
		return
	}

	if id.ID == 0 {
		httputils.HandleError(&w, 500, "Bad Request", "ID not provided", nil)
		return
	}

	if storage.Remove(id.ID) {
		httputils.HandleSuccess(&w, entities.ID{ID: id.ID})
	} else {
		httputils.HandleError(&w, 400, "Bad Request", "ID not found", nil)
	}
}
