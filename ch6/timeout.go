package ch6

import (
	"fmt"
	"math/rand"
	"time"
)

// Another popular option (other than buffering, or selecting) is to timeout
// With a timeout, we're willing to block for some time, but not forever.
// This is also something easy to achieve in Go
// The syntax might be hard to follow but it's such a neat and useful feature that I couldn't leave it out

// To block for a maximum amount of time, we can use the `time.After` function.
// Let's look at it then try to peek beyond the magic.
// To use this, our sender becomes:

func Initialize_workers_and_give_them_work_to_do_using_select_and_with_a_timeout() {
	// To use this, the first thing we'd do is start some workers:
	c := make(chan int) // Note, not buffered

	for i := 0; i < 5; i++ {
		worker := &LazyWorker{id: i}
		go worker.process(c)
	}

	for {
		select {
		case c <- rand.Int():
		// `time.After` returns a channel, so we can `select` from it
		// The channel is written to after the specified time expires
		case <-time.After(time.Millisecond * 100):
			fmt.Println("timed out")
		}
		time.Sleep(time.Millisecond * 50)
	}
}

func Initialize_workers_and_give_them_work_to_do_using_select_with_a_default_and_with_a_timeout() {
	// To use this, the first thing we'd do is start some workers:
	c := make(chan int) // Note, not buffered

	for i := 0; i < 5; i++ {
		worker := &LazyWorker{id: i}
		go worker.process(c)
	}

	for {
		select {
		case c <- rand.Int():
		// `time.After` returns a channel, so we can `select` from it
		// The channel is written to after the specified time expires
		case <-time.After(time.Millisecond * 100):
			fmt.Println("timed out")
		}
		time.Sleep(time.Millisecond * 50)
	}
}
