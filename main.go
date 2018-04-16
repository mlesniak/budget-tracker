package main

import (
	"log"
	"os"
)

// TODO ML create generalized abstraction / library? for assets vs local storage for docker
// TODO ML Favicon
// TODO ML README
// TODO ML r/golang posting
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
