package generalutil

import (
	"github.com/joho/godotenv"
	"log"
)

// SetupCheck : Utility to Check .env File
func SetupCheck() bool {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return false
	}
	return true
}
