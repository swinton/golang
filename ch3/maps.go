package ch3

import "fmt"

// Maps in Go are what other languages call hashtables or dictionaries.
// They work as you expect: you define a key and value, and can get, set and delete values from it.

func Make_a_simple_map() map[string]int {
	// Maps, like slices, are created with the `make` function
	lookup := make(map[string]int)
	lookup["goku"] = 9001

	power, exists := lookup["vegeta"]
	fmt.Println(power, exists) // prints 0 false (0 is the default value for an integer)

	power, exists = lookup["goku"]
	fmt.Println(power, exists) // prints 9001 true

	return lookup
}

func Get_the_number_of_keys_in_a_map() {
	lookup := Make_a_simple_map()

	// To get the number of keys, we use `len`
	total := len(lookup)

	fmt.Println(total) // prints 1
}

func Delete_from_a_map() {
	lookup := Make_a_simple_map()

	// has no return, can be called on a non-existing key
	delete(lookup, "goku")

	total := len(lookup)
	fmt.Println(total) // prints 0
}

func Set_an_initial_size_when_making_a_map() map[string]int {
	// Maps grow dynamically. However, we can supply a second argument to `make` to
	// set an initial size:
	lookup := make(map[string]int, 100)
	return lookup
}

// 	Map as field of struct
type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}

func Initialize_struct_having_map_as_field() *Saiyan {
	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}
	goku.Friends["krillin"] = &Saiyan{Name: "Krillin"}

	return goku
}

func Declare_a_composite_literal() map[string]int {
	// This approach is specific to maps and arrays
	lookup := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}

	return lookup
}

func Iterate_over_map_with_for() {
	lookup := Declare_a_composite_literal()

	// We can iterate over a map using a `for` loop combined with the `range` keyword
	// Iteration over maps isn't ordered
	for key, value := range lookup {
		fmt.Println(key, value)
	}
}
