package ch6

import (
	"fmt"
	"math/rand"
	"time"
)

// Given the previous example, what happens if we have more data coming in than we can handle?
// One strategy, it all workers are busy, is to temporarily store the data in some sort of queue.
// Channels have this buffering capability built-in.
// When we created our channel with `make`, we can give our channel a length

// var c = make(chan int, 100) // We can give our channel a length

// Buffered channels merely provide:
// - a queue for pending work
// - a good way to deal with a sudden spike

type LazyWorker struct {
	id int
}

func (w LazyWorker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker id %d, got data %d\n", w.id, data)
		// Artificially slow down each worker so the work builds up
		time.Sleep(time.Millisecond * 500)
	}
}

func Initialize_workers_and_give_them_work_to_do_using_a_buffered_channel() {
	// To use this, the first thing we'd do is start some workers:
	c := make(chan int, 100) // We can give our channel a length

	for i := 0; i < 5; i++ {
		worker := &LazyWorker{id: i}
		go worker.process(c)
	}

	// And then we can give them some work:
	for {
		// what Go guarantees, is that the data we send to a channel will only
		// be received by a single receiver
		c <- rand.Int()

		// Monitor the length of the channel
		fmt.Printf(">>> channel length is %d\n", len(c))
		time.Sleep(time.Millisecond * 50)
	}
}
