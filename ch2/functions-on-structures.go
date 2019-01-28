package ch2

import "fmt"

// We can associate a method with a structure:

// We say the type *Saiyan is the receiver of the Super method
func (s *Saiyan) Super() {
	s.Power += 10000
}

func Example() {
	goku := &Saiyan{"Goku", 9001}
	goku.Super()
	fmt.Println(goku.Power) // will print 19001
}
