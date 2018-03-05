package main

import (
	"math/rand"
	"strconv"
)

func main() {
	// Later, we will call the HTTP handler here, too...
	InitalizeStorage()

	amount := rand.Float64() * 100
	t := NewTransaction("demo", strconv.FormatFloat(amount, 'g', 2, 32))
	Save(t)
}
