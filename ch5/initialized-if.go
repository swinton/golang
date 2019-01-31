package ch5

func process() (int, bool) {
	return 42, true
}

// Contrived example
// The values aren't available outside the if-statement
// they are available inside any `else if` or `else`
func Show_me_initialized_if_example() bool {
	if _, err := process(); err {
		return true
	}

	return false
}
