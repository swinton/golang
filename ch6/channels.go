package ch6

// `channels` aim at making concurrent programming cleaner and less error-prone.

// Channels help make concurrent programming saner by taking shared data out of the picture.
// A channel is a communication pipe between goroutines which is used to pass data.
// In other words, a goroutine that has data can pass it to another goroutine via a channel.
// The result is that, at any point in time, only one goroutine has access to the data.

// A channel, like everything else, has a type.
// This is the type of data that we'll be passing through our channel.

// For example, to create a channel which can be used to pass an integer around, we'd do:

var c = make(chan int) // The type of this channel is `chan int`

// To pass this channel to a function, our signature looks like
func worker(c chan int) {

}

// Channels support two operations: receiving and sending.
// receiving and sending to and from a channel is blocking

// We send to a channel by doing:
// ```
// CHANNEL <- DATA  // SEND! The data flows into the channel
// ```

// We receive from one by doing
// ```
// VAR := <-CHANNEL  // RECEIVE! The data flows out of the channel
// ```
