package ch5

// Functions are first-class types:
type MathFunc func(a int, b int) int

// This "AddFunc" type can be used anywhere:
// - as a field type
// - as a parameter
// - as a return value

// Implementation of AddFunc is handled separately, decoupled from the type definition

// Applies the given MathFunc to ints a and b
func apply(f MathFunc, a int, b int) int {
	return f(a, b)
}

func adder(a int, b int) int {
	return a + b
}

func multiplier(a int, b int) int {
	return a * b
}

func Do_the_math(a int, b int) (int, int) {
	return apply(adder, a, b), apply(multiplier, a, b)
}
