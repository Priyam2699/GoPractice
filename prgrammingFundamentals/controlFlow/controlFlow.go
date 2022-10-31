package main

import "fmt"

func main() {

	// if the variable is declared and initialized in if statements then it limits the scope
	if x := 42; x != 2 {
		fmt.Println("001")
	}

	fmt.Println("Hello")
	fmt.Println("Priyam")
	// the below line will not run
	//fmt.Println(x)
}
