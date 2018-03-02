package main

import (
	"testing"
)

func TestBalance(t *testing.T) {
	t1 := NewTransaction("income", "10.0")
	t2 := NewTransaction("expense", "-10.0")
	ta := make([]Transaction, 2)
	d := append(ta, t1, t2)

	b := ComputeBalance(d)
	if b.Amount.StringFixed(2) != "0.00" {
		t.Error(b.Amount)
	}
}
