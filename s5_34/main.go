package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "This is from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server 2"
	}
}

func main() {
	fmt.Printf("Select with channels\n")
	fmt.Printf("--------------------\n")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		select {
		case s1 := <-channel1:
			fmt.Printf("case one: %s\n", s1)
		case s2 := <-channel1:
			fmt.Printf("case two: %s\n", s2)
		case s3 := <-channel2:
			fmt.Printf("case three: %s\n", s3)
		case s4 := <-channel2:
			fmt.Printf("case four: %s\n", s4)
			// default:
			// avoiding deadlock
		}
	}
}
