package main

import (
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect len of order finished; expected 5, got %d", len(orderFinished))
		}
	}
}

func TestDineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quartr second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, e := range theTests {
		orderFinished = []string{}

		eatTime = e.delay
		sleepTime = e.delay
		thinkTime = e.delay

		dine()
		if len(orderFinished) != 5 {
			t.Errorf("%s: incorrect len of order finished; expected 5, got %d",
				e.name, len(orderFinished))
		}

	}
}
