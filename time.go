package main

import (
	"time"
	"log"
)

// Time is a function which returns the "current" time. Its value is temporarily overwritten
// in unit tests to return a pre-defined time.
var Time = func() time.Time {
	return time.Now()
}

// RestoreTime restores the time retrieval function to a previous value.
//
// Used in unit tests with
//
//     defer RestoreTime(Time)
func RestoreTime(oldTime func() time.Time) {
	log.Println("Restoring time")
	Time = oldTime
}

// MockTime allows to return a fixed value on a call to Time().
//
// Used in unit tests with
//
//     MockTime("2018-03-01 00:00:00")
func MockTime(now string) {
	log.Println("Mocking time to", now)
	Time = func() time.Time {
		mockTime, err := time.Parse("2006-01-02 15:04:05", now)
		if err != nil {
			log.Fatal("Unable to parse time", err)
		}
		return mockTime
	}
}
