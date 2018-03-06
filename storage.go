package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	// This is the usual way to include an SQL driver in golang. Actually we are not using
	// any imports from the package explictly.
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

// Currently, we only support a single user.
const userID = 1

// InitalizeStorage creates the database and creates non-existing tables.
//
// Uses fileName to define the database main file
func InitalizeStorage(fileName string) {
	database, _ = sql.Open("sqlite3", fileName)
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

// Save a new transaction.
func Save(t Transaction) {
	statement, _ := database.Prepare(
		"INSERT INTO transactions (userid, year, month, timestamp, category, amount)" +
			"VALUES (?,?,?,?,?,?)")
	amount, _ := t.Amount.Float64()
	year, month, _ := t.Timestamp.Date()
	statement.Exec(userID, year, month, t.Timestamp, t.Category, amount)
}

// Load all transactions for a given year and month.
//
// Note that currently solely the default user is chosen.
func Load(year, month int) Transactions {
	ts := make(Transactions, 0, 16)

	rows, _ := database.Query(
		"SELECT timestamp, category, amount FROM transactions WHERE "+
			"userid = ? AND year = ? AND month = ?", userID, year, month)

	for rows.Next() {
		var timestamp time.Time
		var category string
		var amount Amount
		rows.Scan(&timestamp, &category, &amount)
		t := Transaction{category, timestamp, amount}
		ts = append(ts, t)
	}

	return ts
}
