package ch1

// This is a good time to point out that functions can return multiple values.
// Let's look at three functions: one with no return value, one with one return value,
// and one with two return values.

func log(message string) {
}

func add(a int, b int) int {
	// return something so it compiles
	return 42
}

func power(name string) (int, bool) {
	// return something so it compiles
	return 42, false
}
