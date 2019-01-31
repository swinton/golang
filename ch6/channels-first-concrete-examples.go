package ch6

import (
	"fmt"
	"math/rand"
	"time"
)

// Consider a system with incoming data that we want to handle in separate goroutines.
// This is a common requirement.
// If we did our data-intensive processing on the goroutine which accepts the incoming data,
// we'd risk timing out clients.
// First, we'll write our worker.
// This could be a simple function, but I'll make it part of a structure since we haven't seen goroutines used like this before:

type Worker struct {
	id int
}

// Our worker is simple.
// It waits until data is available then "processes" it.
// Dutifully, it does this in a loop, forever waiting for more data to process.
func (w Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker id %d, got data %d\n", w.id, data)
	}
}

func Initialize_workers_and_give_them_work_to_do() {
	// To use this, the first thing we'd do is start some workers:
	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	// And then we can give them some work:
	for {
		// what Go guarantees, is that the data we send to a channel will only
		// be received by a single receiver
		c <- rand.Int()
		time.Sleep(time.Millisecond * 50)
	}
}
