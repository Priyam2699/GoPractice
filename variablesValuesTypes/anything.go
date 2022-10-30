package main

import "fmt"

func main() {
	fmt.Println("Hello Priyam, How are you?")

	foo()

	i := 0
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("End")
}

func foo() {

	fmt.Println("I am in foo")
}
