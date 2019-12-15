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
			addCorsHeader(w)
			w.WriteHeader(http.StatusOK)
		})
	router.HandleFunc("/ping", controller.Ping())
	router.HandleFunc("/heroes", controller.Heroes())
	router.HandleFunc("/chapter", controller.Chapter())
	return router
}

func addCorsHeader(w http.ResponseWriter) {
	headers := w.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
}
