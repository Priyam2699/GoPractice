package main

import "fmt"

func main() {

	switch {

	case false:
		fmt.Println("Condition is false")
		fallthrough
	case true:
		fmt.Println("Condition is true")

	}

}
