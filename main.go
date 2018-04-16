package main

import (
	"log"
	"os"
)

// TODO ML Deployment script with update-on-server functionality
// TODO ML README with build instructions
// TODO ML r/golang posting
// TODO ML Favicon on HomeScreen (https://en.wikipedia.org/wiki/Favicon#Browser_implementation)
// TODO ML create generalized abstraction / library? for assets vs local storage for docker
func main() {
	databasePath := os.Getenv("BUDGET_DATABASE")
	if databasePath == "" {
		databasePath = "./data.db"
	}
	InitalizeStorage(databasePath)

	password, found := os.LookupEnv("BUDGET_PASSWORD")
	if !found {
		log.Println("Cookie password not set in BUDGET_PASSWORD. ABORTING")
		return
	}
	StartServer(password)
}
