package ch2

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

// Note `*Person` is not given an explicit field name, this allows us to
//  access the Person fields and methods *directly*
type SaiyanComposed struct {
	*Person
	Power int
}

func NewSaiyanComposed() *SaiyanComposed {
	goku := &SaiyanComposed{
		Person: &Person{"Goku"},
		Power:  9001,
	}

	// We can call the method attached to Person *directly*
	goku.Introduce()

	// The following also works
	goku.Person.Introduce()

	return goku
}
