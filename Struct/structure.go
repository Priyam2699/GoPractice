package main

import "fmt"

// struct in go
type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person

	LTK bool
}

func main() {

	// p1 of type person
	p1 := person{

		first: "Priyam",
		last:  "Dua",
		age:   29,
	}

	p2 := person{

		first: "Ajay",
		last:  "Maheshwary",
	}

	fmt.Println(p1.age, p2.age)

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},

		LTK: true,
	}

	fmt.Println(sa1, sa1.person, sa1.person.first, sa1.last, sa1.age, sa1.LTK)

}
