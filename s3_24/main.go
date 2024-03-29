package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

// producer is a type for structs that holds two channels: one for pizzas, with all
// information for a given pizza order including whether it was made
// successfully, and another to handle end of processing (when we quit the channel)
type producer struct {
	data chan pizzaOrder
	quit chan chan error
}

// pizzaOrder is a type for structs that describes a given pizza order. It has the order
// number, a message indicating what happened to the order, and a boolean
// indicating if the order was successfully completed.
type pizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

// Close is simply a method of closing the channel when we are done with it (i.e.
// something is pushed to the quit channel)
func (p *producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

// makePizza attempts to make a pizza. We generate a random number from 1-12,
// and put in two cases where we can't make the pizza in time. Otherwise,
// we make the pizza without issue. To make things interesting, each pizza
// will take a different length of time to produce (some pizzas are harder than others).
func makePizza(pizzaNumber int) *pizzaOrder {
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
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready", pizzaNumber)
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

// pizzeria is a goroutine that runs in the background and
// calls makePizza to try to make one order each time it iterates through
// the for loop. It executes until it receives something on the quit
// channel. The quit channel does not receive anything until the consumer
// sends it (when the number of orders is greater than or equal to the
// constant NumberOfPizzas).
func pizzeria(pizzaMaker *producer) {
	// keep track of which pizza we are making
	i := 0

	// run forever or until we received a quit notification
	// try to make pizzas
	for {
		// try to make a pizza
		currentPizza := makePizza(i)
		// decision
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza (we send something to the data channel -- a chan PizzaOrder)
			case pizzaMaker.data <- *currentPizza:
			// we want to quit, so send pizzMaker.quit to the quitChan (a chan error)
			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
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

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing chanel!", err)
			}
		}
	}

	// print out the ending message
}
