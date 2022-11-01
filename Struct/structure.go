package main

import "fmt"

// struct in go
type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{

		first: "Priyam",
		last:  "Dua",
		age:   24,
	}

	// anonymous struct
	x := struct {
		email        string
		mobileNumber int
	}{

		email:        "priyamdua26@gmail.com",
		mobileNumber: 5195621651,
	}

	fmt.Println(p1)
	fmt.Println(x)
}
