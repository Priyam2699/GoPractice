package main

import "fmt"

func main() {

	var x int
	fmt.Scan(&x)
	switch {
	case x == 1:
		fmt.Println("Number 1")
	case x == 2:
		fmt.Println("Number 2")
		fallthrough
	case x == 3:
		fmt.Println("Number 3")
		fallthrough
	case x == 4:
		fmt.Println("Number 4")
		fallthrough
	case x%2 == 0:
		fmt.Println("Even number")
	default:
		fmt.Println("this is default")

	}

	var str string
	fmt.Scan(&str)

	switch str {
	case "Priyam", "Dua":
		fmt.Println("Priyam")
	case "James":
		fmt.Println("James")
	default:
		fmt.Println("Name does not match")

	}
}
