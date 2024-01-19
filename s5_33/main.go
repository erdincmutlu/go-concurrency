package main

import (
	"fmt"
	"strings"
)

func shout(ping chan string, pong chan string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}
func main() {
	// create two channels
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Printf("Type something and press ENTER (enter Q to quit)")

	for {
		// print a prompt
		fmt.Printf("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)
		if strings.ToLower(userInput) == "q" {
			break
		}

		ping <- userInput

		// wait for a response
		response := <-pong
		fmt.Printf("Response: %s\n", response)

	}

	fmt.Printf("All done. Closing channels.")
	close(ping)
	close(pong)
}
