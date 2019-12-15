package api

import (
	"github.com/darkarchana/darkarchana-backend/controller"
	"github.com/gorilla/mux"
)

// Routing : list of API routing
func Routing() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", controller.Ping()).Methods("GET")
	router.HandleFunc("/heroes", controller.Heroes()).Methods("GET")
	router.HandleFunc("/chapter", controller.Chapter()).Methods("GET")
	return router
}
