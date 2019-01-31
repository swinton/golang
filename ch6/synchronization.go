package ch6

// Concurrent code needs to be coordinated
// To help with this problem, Go provides `channels`.

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0

// What will the output be?
func Unsynchronized_example() {
	for i := 0; i < 20; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 10)
}

func incr() {
	counter++
	fmt.Println(counter)
}

// The only concurrent thing you can safely do to a variable is to read from it.
// You can have as many **readers** as you want, but **writes need to be synchronized**.
// Commonly, writes are synchronized using a `mutex`

var lock sync.Mutex // default value of a `sync.Mutex` is unlocked
// sync.RWMutex also exists

func Synchronized_example() {
	for i := 0; i < 20; i++ {
		go incr_synchronized()
	}
	time.Sleep(time.Millisecond * 10)
}

func incr_synchronized() {
	// We generally want fine locks; else, we end up with a ten-lane highway that suddenly
	// turns into a one-lane road
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println(counter)
}

// ## DEADLOCKS
//
// if you're using two or more locks around the same code, it's dangerously easy to have
// situations where goroutineA holds lockA but needs access to lockB, while goroutineB holds
// lockB but needs access to lockA

func Deadlock_example() {
	// Acquire, but never release the lock
	go func() { lock.Lock() }()
	time.Sleep(time.Millisecond * 10)
	// fatal error: all goroutines are asleep - deadlock!
	lock.Lock()
}
