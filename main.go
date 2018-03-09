package main

import (
	"log"
	"strconv"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

func main() {
	InitalizeStorage("./data.db")
	startServer()
}

func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/api/transaction/{year}/{month}", transactionHandler)
	port := ":8080"
	log.Println("Starting to listen on port", port)
	http.ListenAndServe(port, r)
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("Transaction handler called. vars=", vars)
	year, err := strconv.Atoi(vars["year"])
	if err != nil {
		log.Println("Unable to parse year", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	month, err := strconv.Atoi(vars["month"])
	if err != nil {
		log.Println("Unable to parse month", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	ts := Load(year, month)
	enc := json.NewEncoder(w)
	enc.Encode(ts)
}
