package main

import (
	"log"
	"net/http"
	"os"

	"github.com/darkarchana/darkarchana-backend/api"
	"github.com/darkarchana/darkarchana-backend/database"
	"github.com/darkarchana/darkarchana-backend/generalutil"
	"github.com/gorilla/handlers"
)

func main() {

	if generalutil.SetupCheck() {
		headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
		originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
		methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

		mux := api.Routing()
		port := os.Getenv("PORT")
		if port == "" {
			port = "3000"
			log.Print("INFO: No PORT environment variable detected, defaulting to " + port)
		}

		database.MongoDbSetup()
		log.Println("Listening on localhost:" + port)
		http.ListenAndServe(":"+port, handlers.CORS(headersOk, originsOk, methodsOk)(mux))
	}
}
