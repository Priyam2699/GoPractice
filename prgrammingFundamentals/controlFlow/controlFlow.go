package main

import "fmt"

func main() {

	var x int
	fmt.Scan(&x)

	if x%2 == 0 && x%3 == 0 {
		fmt.Println("Number divided by 2 and 3")
	} else {
		fmt.Println("Number not divided by 2 and 3")
	}
}
