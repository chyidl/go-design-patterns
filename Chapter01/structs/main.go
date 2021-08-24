package main

import "fmt"

type Person struct {
	// public fields
	Name    string
	Surname string
	Hobbies []string
	// private fields
	id string
}

// struct function on pointer
func (p *Person) GetFullName() string {

	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}

// struct function on value
// full copy was passed
func (p Person) GetHobbies() []string {
	// non-pointer valueMethod when you're returning a immutable private property
	return p.Hobbies
}

func main() {
	p := Person{
		Name:    "Mario",
		Surname: "Castro",
		Hobbies: []string{"cycling", "electronics", "planes"},
		id:      "sa3-223-asd",
	}

	fmt.Printf("%s likes %s, %s and %s\n", p.GetFullName(), p.GetHobbies()[0], p.GetHobbies()[1], p.GetHobbies()[2])

}
