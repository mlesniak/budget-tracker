package main

import (
	"log"

	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	InitalizeStorage("./data.db")
	startServer()
}

func startServer() {
	r := mux.NewRouter()
    r.HandleFunc("/", homeHandler)
	port := ":8080"
	log.Println("Starting to listen on port", port)
	http.ListenAndServe(port, r)
}

func homeHandler(rw http.ResponseWriter,req  *http.Request) {
	bs := []byte("Hello, world")
	rw.Write(bs)
}
