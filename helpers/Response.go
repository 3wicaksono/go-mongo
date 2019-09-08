package helpers

import (
	"encoding/json"
	"net/http"
)

// Response - response json output to client
func Response(w http.ResponseWriter, httpStatus int, data interface{}) {
	if httpStatus == 0 {
		httpStatus = http.StatusOK
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(data)
}
