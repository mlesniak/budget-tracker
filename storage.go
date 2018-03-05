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

// Currently, we only support a single user.
const userId = 1

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

// Save a new transaction.
func Save(t Transaction) {
	statement, _ := database.Prepare(
		"INSERT INTO transactions (userid, year, month, timestamp, category, amount)" +
			"VALUES (?,?,?,?,?,?)")
	amount, _ := t.Amount.Float64()
	year, month, _ := t.Timestamp.Date()
	statement.Exec(userId, year, month, t.Timestamp, t.Category, amount)
}

func Load(year, month int) Transactions {
	fmt.Println("Loading for", year, month)
	return Transactions{}
}

// func fill(database *sql.DB) {
// 	statement, _ := database.Prepare("INSERT INTO storage (key, value) VALUES (?, ?)")
// 	statement.Exec("timestamp", time.Now())
// }

// func query(database *sql.DB) {
// 	rows, _ := database.Query("SELECT id, key, value FROM storage")
// 	var id int
// 	var key string
// 	var value string
// 	for rows.Next() {
// 		rows.Scan(&id, &key, &value)
// 		log.Println(strconv.Itoa(id) + ":" + key + ":" + value)
// 	}
// }
