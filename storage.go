package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	// This is the usual way to include an SQL driver in golang. Actually we are not using
	// any imports from the package explictly.
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func InitalizeStorage() {
	database, _ = sql.Open("sqlite3", "./demo.db")
	executeFile("init.sql")
}

func executeFile(fileName string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("No initialization file found.")
		return
	}

	cmds := strings.Split(string(bytes), ";")
	for _, cmd := range cmds {
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		statement, err := database.Prepare(cmd)
		if err != nil {
			fmt.Println("Unable to execute statement: ", cmd)
			return
		}
		statement.Exec()
	}
}

func Save(t Transaction) {
	fmt.Println("Saving:", t) 
}

func Load(year, month int) Transactions {
	fmt.Println("Loading for", year, month)
	return Transactions{}
}