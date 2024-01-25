package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestUpdateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("Hello Erdinc")
	wg.Wait()

	if !strings.Contains(msg, "Hello Erdinc") {
		t.Errorf("Expecting Hello Erdinc in the output")
	}
}

func TestPrintMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "Hello Erdinc"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "Hello Erdinc") {
		t.Errorf("Expecting Hello Erdinc in the output")
	}
}


func TestMain(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Expecting Hello, universe! in the output")
	}

    if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("Expecting Hello, cosmos! in the output")
	}

    if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expecting Hello, world! in the output")
	}

}


