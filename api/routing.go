package api

import (
	"net/http"

	"github.com/darkarchana/darkarchana-backend/controller"
	"github.com/gorilla/mux"
)

// Routing : list of API routing
func Routing() *mux.Router {
	router := mux.NewRouter()
	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Connection", "keep-alive")
			w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
			w.Header().Add("Access-Control-Allow-Headers", "content-type")
			w.Header().Add("Access-Control-Max-Age", "86400")
		})
	router.HandleFunc("/ping", controller.Ping())
	router.HandleFunc("/heroes", controller.Heroes())
	router.HandleFunc("/chapter", controller.Chapter())
	return router
}
