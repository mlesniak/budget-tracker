package main

import (
	"fmt"
	"testing"
)

func TestSave(t *testing.T) {
	defer RestoreTime(Time)
	MockTime("2018-03-01 00:00:00")
	tn := NewTransaction("income", "1234")
	Save(tn)

	tns := Load(2018, 3)
	fmt.Println(tns)
}