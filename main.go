package main

import (
	"log"
	"net/http"

	"github.com/darkarchana/darkarchana-backend/api"
	"github.com/darkarchana/darkarchana-backend/database"
	"github.com/darkarchana/darkarchana-backend/generalutil"
)

func main() {

	if generalutil.SetupCheck() {
		mux := api.Routing()
		database.MongoDbSetup()
		log.Println("Listening on localhost:3000")
		http.ListenAndServe("localhost:3000", mux)
	}
}
