package httphandlers

import (
	"net/http"

	"example.com/service/product/httphandlers/httputils"
	"example.com/service/product/storage"
)

// List returns a list of all products
func List(w http.ResponseWriter, r *http.Request) {
	httputils.HandleSuccess(&w, storage.Get())
}
