package ch3

import "fmt"

// Arrays are efficient but rigid.
// We often don't know the number of elements we'll be dealing with upfront.
// For this, we turn to slices.

// A slice is a lightweight structure that wraps and represents a portion of an array.

func declare_and_initialize_a_slice() []int {
	// a slight variation on how we created an array
	// but our slice isn't declared with a length
	scores := []int{1, 4, 293, 4, 9}
	return scores
}

func declare_a_slice_using_make() []int {
	// `make` does more than `new`
	// as well as allocating memory, `make` also initializes values

	// Here we initialize a slice with a length (size of the slice) of 10
	// and a capacity of 10 (size of the underlying array)
	scores := make([]int, 10)

	// The two (length, capacity) can be specified separately:
	// Here, a slice with a length of 0 but with a capacity of 10
	scores = make([]int, 0, 10)

	return scores
}

func length_less_than_capacity_fail() {
	// FAILs, we need to explicitly expand our slice in order to access those elements
	scores := make([]int, 0, 10)
	scores[7] = 9033
	fmt.Println(scores)
}

func length_less_than_capacity_succeed() {
	// One way to expand a slice is via `append`:
	scores := make([]int, 0, 10)
	scores = append(scores, 5)
	fmt.Println(scores) // prints [5]
}

func Reslice_our_slice() {
	// So that we can set the element at position 7
	scores := make([]int, 0, 10)
	// Reslice the slice
	scores = scores[0:8]
	// Set the element at position 7
	scores[7] = 9033
	fmt.Println(scores) // prints [0 0 0 0 0 0 0 9033]
}

// How large can we resize a slice? Up to its capacity which, in this case, is 10.

// We can use `append` to overcome this

// If the underlying array is full, `append` will create a new larger array and copy the values over

func Go_grows_arrays_with_a_2x_algorithm() {
	// can you guess what the following will output?

	// scores is a slice of length 0, with an underlying array of capacity 5
	scores := make([]int, 0, 5)
	c := cap(scores)

	// The cap built-in function returns the capacity of v, according to its type
	// Slice: the maximum length the slice can reach when resliced;
	fmt.Println(c) // prints 5

	for i := 0; i < 25; i++ {
		scores = append(scores, i)

		// if our capacity has changed,
		// Go had to grow our array to accommodate the new data
		if cap(scores) != c {
			// The capacity will double each time
			// So we should see: 5, 10, 20, 40
			c = cap(scores)
			fmt.Println(c)
		}
	}
}

func Values_are_appended_at_the_end_of_an_initialized_slice() {
	// A slice of length 5, with underlying capacity 5
	scores := make([]int, 5)
	fmt.Println(scores) // [0, 0, 0, 0, 0, 0]
	scores = append(scores, 9332)
	fmt.Println(scores) // [0, 0, 0, 0, 0, 9332]
}

func The_four_ways_to_initialize_a_slice() {
	// You use this when you know the values that you want in the array ahead of time...
	known_names := []string{"john", "paul", "ringo", "george"}
	fmt.Println(known_names) // [john paul ringo george]

	// useful when you'll be writing into specific indexes of a slice
	checks := make([]bool, 10)
	fmt.Println(checks) // [false false false false false false false false false false]

	// a nil slice and is used in conjunction with `append`, when the number of elements is unknown
	var unknown_names []string
	fmt.Println(unknown_names) // []

	scores := make([]int, 0, 20)
	fmt.Println(scores) // []
}
