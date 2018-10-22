package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"mime"

	"github.com/gorilla/mux"
)

type httpHandler = func(w http.ResponseWriter, r *http.Request)

// Password used for authentication layer
var password string

// StartServer starts the HTTP server and configures all necessary routes.
func StartServer(pw string) {
	password = pw

	// TODO ML Refactor into multiple files / restructure this file.
	r := mux.NewRouter()
	r.HandleFunc("/api/transaction/{year}/{month}", requireAuthentication(listHandler))
	r.HandleFunc("/api/authenticated", requireAuthentication(authHandler))
	r.HandleFunc("/api/transaction/{year}/{month}/budget", requireAuthentication(budgetHandler))
	r.HandleFunc("/api/transaction", requireAuthentication(postHandler)).
		Methods("POST")
	r.PathPrefix("/").HandlerFunc(fileHandler)
	port := "8080"
	log.Println("Starting to listen on port", port)
	
	// http.ListenAndServe(port, r)
	err := http.ListenAndServeTLS(":" + port, "server.crt", "server.key", r)
    if err != nil {
        log.Fatal("Unable to start listening for connections ", err)
    }
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat("static/")
	localMode := true
	if err != nil {
		localMode = false
	}

	if localMode {
		f := http.FileServer(http.Dir("./static/"))
		f.ServeHTTP(w, r)
		return
	}

	// Serving from files in go-bindata only.
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	if upath == "/" {
		upath = "/index.html"
	}
	bytes, err := Asset("static" + upath)
	if err != nil {
		log.Println("File not found", upath)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Handle Content-Type: Header
	nameParts := strings.Split(upath, ".")
	ext := ""
	if len(nameParts) > 0 {
		ext = nameParts[len(nameParts) - 1]
	}
	contentType := mime.TypeByExtension("." + ext)
	if contentType != "" {
		w.Header().Add("Content-Type", contentType)
	}

	w.Write(bytes)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List transactions handler called. vars=", mux.Vars(r))
	year, month, err := parseStandardFields(w, r)
	if err != nil {
		return
	}

	w.Header()["Content-Type"] = []string{"application/json"}
	ts := Load(year, month)
	enc := json.NewEncoder(w)
	enc.Encode(ts)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	bs := []byte("OK")
	w.Write(bs)
}

func budgetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Budget transactions handler called. vars=", mux.Vars(r))
	year, month, err := parseStandardFields(w, r)
	if err != nil {
		return
	}

	w.Header()["Content-Type"] = []string{"application/json"}
	ts := Load(year, month)
	budget := ComputeBudget(ts)
	enc := json.NewEncoder(w)
	enc.Encode(budget)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("Post transaction handler called. vars=", vars)

	dec := json.NewDecoder(r.Body)
	var t Transaction
	dec.Decode(&t)
	w.WriteHeader(http.StatusOK)
	trans := NewTransaction(t.Description, t.Amount.StringFixed(2))
	Save(trans)
}

func parseStandardFields(w http.ResponseWriter, r *http.Request) (int, int, error) {
	vars := mux.Vars(r)
	year, err := strconv.Atoi(vars["year"])
	if err != nil {
		log.Println("Unable to parse year", err)
		w.WriteHeader(http.StatusBadRequest)
		return 0, 0, err
	}
	month, err := strconv.Atoi(vars["month"])
	if err != nil {
		log.Println("Unable to parse month", err)
		w.WriteHeader(http.StatusBadRequest)
		return 0, 0, err
	}

	return year, month, nil
}

func requireAuthentication(fn httpHandler) httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isAuthenticated(r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fn(w, r)
	}
}

func isAuthenticated(r *http.Request) bool {
	cookie, err := r.Cookie("auth")
	if err != nil || cookie.Value != password {
		log.Println("Cookie password not set or invalid")
		return false
	}

	return true
}
