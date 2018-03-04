package main

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

// A Transaction describes a single income or expense.
type Transaction struct {
	Category  string
	Timestamp time.Time
	Amount    Amount
}

// An Amount describes a monetary value.
type Amount = decimal.Decimal

// Transactions encapsulate a list of single transactions.
type Transactions = []Transaction

func (t Transaction) String() string {
	return fmt.Sprintf("(%v) %-10v %v",
		t.Timestamp.Format("2006-01-02 15:04:05"), t.Category, t.Amount)
}

// NewTransaction creates a new transaction with the current timestamp.
func NewTransaction(category string, amount string) Transaction {
	d, _ := decimal.NewFromString(amount)
	return Transaction{category, Time(), d}
}

// ComputeBalance computes the overall balance over all transactions.
func ComputeBalance(transactions Transactions) Amount {
	amount := decimal.NewFromFloat(0)
	for _, t := range transactions {
		amount = amount.Add(t.Amount)
	}
	return amount
}

// ComputeBudget computes the remaining budget for the whole month and per day.
//
// The returned tuple contains (month, day) budget.
func ComputeBudget(transactions Transactions) (Amount, Amount) {
	balance := ComputeBalance(transactions)
	remainingDays := getRemainingDays()
	dailyBudget := computeDailyBudget(balance, remainingDays)
	return balance, dailyBudget
}

// getRemainingDays computes the number of remaining days in the current month.
func getRemainingDays() int {
	// Ignore leap years for now.
	now := Time()
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	year, month, day := now.Date()
	location := now.Location()
	_, _, today := time.Date(year, month, day, 0, 0, 0, 0, location).Date()
	lastDay := days[month-1]
	remainingDays := lastDay - today + 1
	return remainingDays
}

func computeDailyBudget(balance Amount, days int) Amount {
	daysDecimal := decimal.NewFromFloat(float64(days))
	dailyAmount := balance.Div(daysDecimal)
	return dailyAmount
}