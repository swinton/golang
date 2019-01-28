package ch2

import "fmt"

// Go doesn't support overloading

// However, because implicit composition is really just a compiler trick,
// we can "overwrite" the functions of a composed type.
// For example, our `Saiyan` structure can have its own `Introduce` function:

type Human struct {
	Name string
}

func (p *Human) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type SaiyanOverloaded struct {
	*Human
	Power int
}

func (s *SaiyanOverloaded) Introduce() {
	fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}

func NewSaiyanOverloaded() *SaiyanOverloaded {
	goku := &SaiyanOverloaded{
		Human: &Human{"Goku"},
		Power: 9000,
	}

	// We can call the overloaded method attached to Person *directly*
	goku.Introduce()

	// We can still call the original method
	goku.Human.Introduce()

	return goku
}
