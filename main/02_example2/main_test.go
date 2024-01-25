package main

import "testing"

func TestUpdateMessage(t *testing.T) {
	msg = "Hello, world!"

	wg.Add(2)
	updateMessage("x")
	updateMessage("Goodbye, cruel world!")
	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Errorf("incorrect value in msg")
	}
}
