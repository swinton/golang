package ch1

import (
	"fmt"
)

func long_declaration() {
	// The most explicit way to deal with variable declaration
	// and assignment in Go is also the most verbose:
	var power int
	power = 9000
	fmt.Println("It's over", power)
}

func short_declaration() {
	// `:=` is used to declare the variable as well as assign a value to it
	power := 9000
	fmt.Println("It's over", power)
}

func Variables_and_declarations() {
	long_declaration()
	short_declaration()
}
