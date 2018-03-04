package main

import (
	"fmt"
)

func Save(t Transaction) {
	fmt.Println("Saving:", t) 
}

func Load(year, month int) Transactions {
	fmt.Println("Loading for", year, month)
	return Transactions{}
}