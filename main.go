package main

import (
	"log"
	"net/http"
	"os"

	"github.com/darkarchana/darkarchana-backend/api"
	"github.com/darkarchana/darkarchana-backend/database"
	"github.com/darkarchana/darkarchana-backend/generalutil"
)

func main() {

	if generalutil.SetupCheck() {
		mux := api.Routing()
		port := os.Getenv("PORT")
		if port == "" {
			port = "3000"
			log.Print("INFO: No PORT environment variable detected, defaulting to " + port)
		}

		database.MongoDbSetup()
		log.Println("Listening on localhost:" + port)
		http.ListenAndServe(":"+port, mux)
	}
}
