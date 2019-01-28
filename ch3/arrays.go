package ch3

import "fmt"

// In Go, arrays are fixed, not dynamic
// Declaring an array requires that we specify the size, and once
//  the size is specified, it cannot grow

func declare_an_array() [10]int {
	// The above array can hold up to 10 scores using indexes `scores[0]` through `scores[9]`.
	// Attempts to access an out of range index in the array will result in a compiler or runtime error.
	var scores [10]int
	scores[0] = 339

	return scores
}

func declare_and_initialize_an_array() [4]int {
	// We can initialize the array with values:
	scores := [4]int{9001, 9333, 212, 33}
	return scores
}

func Loop_over_array_items() {
	// We can use `len` to get the length of the array. `range` can be used to iterate over it:
	scores := declare_and_initialize_an_array()
	for index, value := range scores {
		fmt.Println(index, ":", value)
	}
}
