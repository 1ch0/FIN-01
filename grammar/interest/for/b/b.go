package b

import "fmt"

type R struct {
	name string
}

func (r *R) SetupWithManager(name string) error {
	fmt.Printf("b: %s\n", name)
	return nil
}

func Setup(name string) error {
	r := R{
		name: name,
	}

	return r.SetupWithManager(name)
}
