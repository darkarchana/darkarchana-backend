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
		credentialOk := handlers.AllowCredentials()
		exposedHeadersOk := handlers.ExposedHeaders([]string{"*"})
		headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
		originsOk := handlers.AllowedOrigins([]string{"*"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})

		mux := api.Routing()
		port := os.Getenv("PORT")
		if port == "" {
			port = "3000"
			log.Print("INFO: No PORT environment variable detected, defaulting to " + port)
		}

		database.MongoDbSetup()
		log.Println("Listening on localhost:" + port)
		http.ListenAndServe(":"+port, handlers.CORS(credentialOk, exposedHeadersOk, headersOk, originsOk, methodsOk)(mux))
	}
}
