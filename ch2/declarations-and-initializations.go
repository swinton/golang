package ch2

// Simple structure:

type Saiyan struct {
	Name  string
	Power int
}

// When we first looked at variables and declarations, we looked only at built-in types,
// like integers and strings. Now that we're talking about structures,
// we need to expand that conversation to include pointers.

// The simplest way to create a value of our structure is:

func simple_declaration_example() Saiyan {
	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}

	return goku
}

func declare_and_initialize_separately() Saiyan {
	goku := Saiyan{}
	goku.Name = "Goku"
	goku.Power = 9000

	return goku
}

// Our function returns an 'address of type Saiyan'
func return_a_pointer() *Saiyan {
	// Use the 'address of' operator (&) to get the address of our value
	goku := &Saiyan{Name: "Goku", Power: 9000}
	return goku
}
