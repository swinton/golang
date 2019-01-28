package ch2

// Fields can be of any type -- including other structures and types
// that we haven't explored yet such as arrays, maps, interfaces and functions.

type SaiyanExpanded struct {
	Name   string
	Power  int
	Father *SaiyanExpanded
}

func NewSaiyanExpanded() *SaiyanExpanded {
	gohan := &SaiyanExpanded{
		Name:  "Gohan",
		Power: 1000,
		Father: &SaiyanExpanded{
			Name:   "Goku",
			Power:  9001,
			Father: nil,
		},
	}

	return gohan
}
