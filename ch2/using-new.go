package ch2

// `new` is used to allocate the memory required by a type

func NewExample() *Saiyan {
	// same as
	// goku := &Saiyan{}
	goku := new(Saiyan) // returns a pointer to a Saiyan
	return goku
}
