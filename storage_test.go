package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitalizeStorage("./_test.db")
	code := m.Run()
	// defer ... will not work here since os.Exit() terminates everything.
	os.Remove("./_test.db")
	os.Exit(code)
}

func TestSave(t *testing.T) {
	defer RestoreTime(Time)
	MockTime("2018-03-02 00:00:00")
	Save(NewTransaction("income2", "5678"))
	MockTime("2018-03-01 00:00:00")
	Save(NewTransaction("income1", "1234"))
	
	tns := Load(2018, 3)
	for _, t := range tns {
		fmt.Println(t)
	}
	if len(tns) != 2 {
		t.Fatal("Wrong number of transactions")
	}
	if tns[0].Description != "income1"  {
		t.Fatal("income1 not first")
	}
	if tns[1].Description != "income2"  {
		t.Fatal("income2 not second")
	}
}
