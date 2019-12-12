package api

import (
	"github.com/darkarchana/darkarchana-backend/controller"
	"net/http"
)

// Routing : list of API routing
func Routing() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", controller.Ping())
	mux.HandleFunc("/heroes", controller.Heroes())
	return mux
}
