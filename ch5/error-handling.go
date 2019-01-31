package ch5

// Go's preferred way to deal with errors is through return values, not exceptions

// Consider strconv

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Consider_the_strconv_way() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	// takes a string and tries to convert it to an integer
	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("not a valid number")
	} else {
		fmt.Println(n)
	}
}

// Creating our own errors...
// by importing the `errors` package and using it in the `New` function

func Gimme_an_error(count int) error {
	if count < 1 {
		return errors.New("Invalid count")
	}
	return nil
}

// Error variables a common pattern, e.g. from the `io` package we have:
var EOF = errors.New("EOF") // this is a package variable, defined outside a function
