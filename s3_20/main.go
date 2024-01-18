package main

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

func main() {
	// seed the random number generator

	// print out a message

	// create a producer

	// run the producer in the background

	// create and run consumer

	// print out the ending message
}
