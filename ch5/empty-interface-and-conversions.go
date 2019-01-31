package ch5

import "fmt"

// Go has no built-in "base class"
// What it does have is a built-in interface, with no fields or methods:
// - `interface{}`
// Every type fulfills the contract of the empty interface

// This is valid...
func Add(a interface{}, b interface{}) interface{} {
	return nil
}

// To convert an interface variable to an explicit type, you use `.(TYPE)`:
func Convert_and_then_add(a interface{}, b interface{}) interface{} {
	return a.(int) + b.(int)
}

func Convert_with_switch(a interface{}, b interface{}) interface{} {
	switch a.(type) {
	case int:
		fmt.Printf("a is now an int and equals %d\n", a)
	default:
		// ...
	}
	return nil
}
