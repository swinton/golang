package ch5

import "fmt"

func Convert_string_to_array_of_bytes(s string) []byte {
	// This creates a copy of the data
	return []byte(s)
}

func Convert_array_of_bytes_to_string(b []byte) string {
	// This creates a copy of the data
	return string(b)
}

func Strings_are_made_of_runes() int {
	return len("椒") // returns 3
}

func Iterating_over_a_string_produces_runes() {
	s := "椒"
	// Produces only 1 value, despite the length being 3
	// When you turn a string into a `[]byte` you'll get each individual byte
	for index, value := range s {
		fmt.Println(index, ":", value)
	}
}
