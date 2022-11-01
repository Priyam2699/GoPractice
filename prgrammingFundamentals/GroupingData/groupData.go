package main

import "fmt"

func main() {

	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[1:])  // start from position 1
	fmt.Println(x[1:4]) // start from 1 and end at position 4
	fmt.Println(x[:4])  // end at position 4

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
