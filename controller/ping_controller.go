package controller

import (
	"encoding/json"
	"github.com/darkarchana/darkarchana-backend/view"
	"net/http"
)

// Ping : API Ping
func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := view.Status{
				Code:     http.StatusOK,
				Response: "Server is Alive",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
