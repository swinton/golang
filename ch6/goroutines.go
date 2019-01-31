package ch6

// A goroutine is similar to a thread, but it is scheduled by Go, not the OS.
// Code that runs in a goroutine can run concurrently with other code
// Multiple goroutines will end up running on the same underlying OS thread
// On modern hardware, it's possible to have millions of goroutines

import (
	"fmt"
	"time"
)

func process() {
	fmt.Println("processing")
}

func Goroutine_example() {
	fmt.Println("start")

	// To start a goroutine, we simply use the `go` keyword followed by the function we want to execute
	go process()

	// We sleep because the main process exits before the goroutine gets a chance to execute
	time.Sleep(time.Millisecond * 1000) // this is bad, don't do this!
	fmt.Println("done")
}

func Goroutine_anonymous_function_example() {
	fmt.Println("start")

	// If we just want to run a bit of code, such as the above, we can use an anonymous function.
	// Do note that anonymous functions aren't only used with goroutines, however
	go func() {
		fmt.Println("processing")
	}()

	// We sleep because the main process exits before the goroutine gets a chance to execute
	time.Sleep(time.Millisecond * 10) // this is bad, don't do this!
	fmt.Println("done")
}
