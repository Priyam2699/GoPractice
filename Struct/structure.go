package main

import "fmt"

// struct in go
type person struct {
	first string
	last  string
	age   int
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

}
