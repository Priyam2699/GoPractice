package main

import "fmt"

func main() {

	var x int
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println("Even")
	} else if x < 0 {
		fmt.Println("Number if less than 0")
	} else {
		fmt.Println("Number is odd")
	}

}
