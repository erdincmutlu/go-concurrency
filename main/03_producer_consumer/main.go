package main

import "github.com/fatih/color"

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

func pizzeria(pizzaMaker *producer) {
	// keep track of which pizza we are making

	// run forever or until we received a quit notification
	// try to make pizzas
	for {
		// try to make a pizza
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
