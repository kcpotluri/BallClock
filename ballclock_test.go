package main

import (
	"testing"
)

func Test_Clock(t *testing.T) {
	cl := clock{}
	cl.init(30)
	cl.startClock()
	if cl.noOfHalfDays/2 != 15 {
		t.Fatal("Failed to get the 15 days for 30 balls")
	}
}

func Test_MinClock(t *testing.T) {
	minClock := clock{}
	minClock.init(30)
	minClock.runClockWithMin(325)
	if len(minClock.mins) != 0 {
		t.Fatal("Expected to get empty queue for min after iteration for 30 ball and 325 min ")
	}
}
