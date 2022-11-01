package main

import "fmt"

func main() {

	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
	fmt.Println(x)

	// appends element at end of slice

	x = append(x, 85, 23, 76)

	fmt.Println(len(x))
	fmt.Println(x)

	y := []int{84, 36, 14, 97}

	// appending a slice to other slice
	x = append(x, y...)
	fmt.Println(len(x))
	fmt.Println(x)

}
