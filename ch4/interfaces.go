package ch4

import "fmt"

// Interfaces are types that define a contract but not an implementation.
// They help decouple your code from specific implementations

type Logger interface {
	Log(message string)
}

// Interfaces are used like any other type...

// ... As a field for a structure
type Server struct {
	logger Logger
}

// ... or a function parameter (or return value):
func process(logger Logger) {
	logger.Log("hello!")
}

// Implementation of an interface is *implicit* in Go

// Here, `ConsoleLogger` can be used wherever a type Logger is required, as it
// implements the Log(message string) "interface"
type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

// The standard library is full of interfaces
// The `io` package has a handful of popular ones such as `io.Reader`, `io.Writer`, and `io.Closer`

// Interfaces can also compose other interfaces
