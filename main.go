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
	Amount    decimal.Decimal
}

func (t Transaction) String() string {
	return fmt.Sprintf("(%v) %-10v %v",
		t.Timestamp.Format("2006-01-02 15:04:05"), t.Category, t.Amount)
}

// Add two Transaction together.
// Note that the category is empty after that.
func (t Transaction) Add(t2 Transaction) Transaction {
	tn := Transaction{}
	tn.Timestamp = time.Now()
	tn.Amount = t.Amount.Add(t2.Amount)
	return tn
}

// NewTransaction creates a new transaction with the current timestamp.
func NewTransaction(category string, amount string) Transaction {
	d, _ := decimal.NewFromString(amount)
	return Transaction{category, time.Now(), d}
}

func main() {
	t1 := Transaction{"income", time.Now(), decimal.NewFromFloat(10.0)}
	t2 := Transaction{"income", time.Now(), decimal.NewFromFloat(20.3)}
	t3 := t1.Add(t2)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
	t4 := Transaction{"expense", time.Now(), decimal.NewFromFloat(-40.12)}
	t5 := t3.Add(t4)
	fmt.Println(t5)
}
