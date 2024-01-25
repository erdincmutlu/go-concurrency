package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type producer struct {
	data chan pizzaOrder
	quit chan chan error
}

type pizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) pizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order %d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)
		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizaa #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizzaorder #%d is ready", pizzaNumber)
		}

		p := pizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}

	return &pizzaOrder{
		pizzaNumber: pizzaNumber,
	}

}

func pizzeria(pizzaMaker *producer) {
	// keep track of which pizza we are making
	i := 0

	// run forever or until we received a quit notification
	// try to make pizzas
	for {
		// try to make a pizza
		currentPizza := makePizza(i)
		// decision
	}
}

func main() {
	// seed the random number generator
	// Not needed from Go 1.20

	// print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &producer{
		data: make(chan pizzaOrder),
		quit: make(chan chan error),
	}

	go pizzeria(pizzaJob)
	// run the producer in the background

	// create and run consumer

	// print out the ending message
}
