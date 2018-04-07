package main

import (
	"os"
	"log")

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
