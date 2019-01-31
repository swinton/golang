package ch6

import (
	"fmt"
	"math/rand"
	"time"
)

// Even with buffering, there comes a point where we need to start dropping messages.
// We can't use up an infinite amount of memory hoping a worker frees up.
// For this, we use Go's `select`.

// With `select`, we can provide code for when the channel isn't available to send to.

func Initialize_workers_and_give_them_work_to_do_using_select() {
	// To use this, the first thing we'd do is start some workers:
	c := make(chan int) // Note, not buffered

	for i := 0; i < 5; i++ {
		worker := &LazyWorker{id: i}
		go worker.process(c)
	}

	for {
		select {
		case c <- rand.Int():
			//optional code here
		default:
			//this can be left empty to silently drop the data
			fmt.Println("dropped")
		}
		time.Sleep(time.Millisecond * 50)
	}
}

// This is only the start of what we can accomplish with `select`.
// A main purpose of select is to manage multiple channels.
// Given multiple channels, `select` will block until the first one becomes available.
// If no channel is available, `default` is executed if one is provided.
// A channel is randomly picked when multiple are available.
