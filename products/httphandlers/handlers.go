package httphandlers

import (
	"log"
	"net/http"

	"example.com/service/product/httphandlers/httputils"
)

// HandleRequest performs method dispatching
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r.Method)
	switch r.Method {
	case http.MethodGet:
		List(w, r)
		break
	case http.MethodPost:
		Add(w, r)
		break
	case http.MethodDelete:
		Remove(w, r)
		break
	default:
		httputils.HandleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
