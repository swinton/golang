package ch2

// Structures don't have constructors.
// Instead, you create a function that returns an instance of the desired type (like a factory):

// Note, we don't *have* to return a pointer
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}
