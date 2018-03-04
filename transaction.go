package main

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

// Time is a function which returns the "current" time. Its value is temporarily overwritten
// in unit tests to return a pre-defined time.
var Time = func() time.Time {
	return time.Now()
}


// A Transaction describes a single income or expense.
type Transaction struct {
	Category  string
	Timestamp time.Time
	Amount    decimal.Decimal
}

// Transactions encapsulate a list of single transactions.
type Transactions = []Transaction

func (t Transaction) String() string {
	return fmt.Sprintf("(%v) %-10v %v",
		t.Timestamp.Format("2006-01-02 15:04:05"), t.Category, t.Amount)
}

// Add two Transaction together.
// Note that the category is empty after that.
func (t Transaction) Add(t2 Transaction) Transaction {
	tn := Transaction{}
	tn.Timestamp = Time()
	tn.Amount = t.Amount.Add(t2.Amount)
	return tn
}

// NewTransaction creates a new transaction with the current timestamp.
func NewTransaction(category string, amount string) Transaction {
	d, _ := decimal.NewFromString(amount)
	return Transaction{category, Time(), d}
}

// ComputeBalance computes the overall balance over all transactions.
func ComputeBalance(transactions Transactions) Transaction {
	tn := Transaction{}
	for _, t := range transactions {
		tn = tn.Add(t)
	}
	return tn
}


// ComputeBudget computes the remaining budget for the whole month and per day.
//
// The returned tuple contains (month, day) budget.
func ComputeBudget(transactions Transactions) (Transaction, Transaction)  {
	// Determine overall balance
	monthlyBudget := ComputeBalance(transactions)

	// Ignore leap years for now.
	days := []int{31,28,31,30,31,30,31,31,30,31,30,31}
	now := Time()
	year, month, day := now.Date()
	location := now.Location()

	_, _, today := time.Date(year, month, day, 0, 0, 0, 0, location).Date()
	lastDay := days[month - 1]
	remainingDays := lastDay - today + 1

	daysDecimal := decimal.NewFromFloat(float64(remainingDays))
	dailyAmount := monthlyBudget.Amount.Div(daysDecimal)

	dailyBudget := Transaction{"", Time(), dailyAmount}

	return monthlyBudget, dailyBudget
}