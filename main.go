package main

import (
	"log"
	"math/rand"
	"time"
	"strconv"
)

func main() {
	// Later, we will call the HTTP handler here, too...
	InitalizeStorage("./data.db")
	
	demo()
}

func demo() {
	log.Println("Starting demo")
	rand.Seed(time.Now().UTC().UnixNano())
	amount := rand.Float64() * 100
	t := NewTransaction("demo", strconv.FormatFloat(amount, 'g', 2, 32))
	Save(t)

	ts := Load(2018, 3)
	for _, t := range ts {
		log.Println("Load", t)
	}
}
