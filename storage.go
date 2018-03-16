package main

import (
	"database/sql"
	"io/ioutil"
	"log"
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
	log.Println("Opening database", fileName)
	var err error
	database, err = sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal("Unable to open database", err)
	}
	executeFile("init.sql")
}

func executeFile(fileName string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("No initialization file found:", fileName)
	}
	log.Println("Executing init script", fileName)

	cmds := strings.Split(string(bytes), ";")
	for _, cmd := range cmds {
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		statement, err := database.Prepare(cmd)
		if err != nil {
			log.Fatal("Unable to execute statement:", cmd)
		}
		statement.Exec()
	}
}

// Save a new transaction.
func Save(t Transaction) {
	statement, err := database.Prepare(
		"INSERT INTO transactions (userid, year, month, timestamp, description, amount)" +
			"VALUES (?,?,?,?,?,?)")
	if err != nil {
		log.Fatal("Invalid insert query", err)
	}
	amount, _ := t.Amount.Float64()
	year, month, _ := t.Timestamp.Date()
	log.Println("Saving transaction", t)
	statement.Exec(userID, year, month, t.Timestamp, t.Description, amount)
}

// Load all transactions for a given year and month.
//
// Note that currently solely the default user is chosen.
func Load(year, month int) Transactions {
	ts := make(Transactions, 0, 16)

	log.Println("Loading transactions for year=", year, "month=", month)
	rows, err := database.Query(
		"SELECT timestamp, description, amount FROM transactions WHERE "+
			"userid = ? AND year = ? AND month = ? ORDER BY timestamp ASC", userID, year, month)
	if err != nil {
		log.Fatal("Unable to execute query", err)
	}
	for rows.Next() {
		var timestamp time.Time
		var description string
		var amount Amount
		rows.Scan(&timestamp, &description, &amount)
		t := Transaction{description, timestamp, amount}
		ts = append(ts, t)
	}

	return ts
}
