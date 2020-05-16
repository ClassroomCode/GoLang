package httputils

import (
	"encoding/json"
	"net/http"
)

// HandleSuccess is a helper function for generating HTTP 200
func HandleSuccess(w *http.ResponseWriter, result interface{}) {
	writer := *w

	marshalled, err := json.Marshal(result)

	if err != nil {
		HandleError(w, 500, "Internal Server Error", "Error marshalling response JSON", err)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write(marshalled)
}
