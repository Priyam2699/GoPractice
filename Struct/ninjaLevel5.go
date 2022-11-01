package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	flavors    []string
}

func main() {

	p1 := person{

		first_name: "Priyam",
		last_name:  "Dua",
		flavors:    []string{"Chocolate", "Vanilla", "Stawberry"},
	}

	p2 := person{
		first_name: "Ajay",
		last_name:  "Maheshwary",
		flavors:    []string{"Blueberry", "Chocochip", "Hazelnut"},
	}

	fmt.Println(p1, "\n", p1.last_name)
	fmt.Println(p2, p2.flavors)

}
