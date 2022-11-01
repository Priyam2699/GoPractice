package main

import "fmt"

func main() {

	// array
	var arr [5]int = [5]int{5, 4, 3, 2, 1}

	// in arrays you cannot go beyond the length of array
	//arr[6] = 0

	fmt.Printf("%v\t%T\n", arr, arr)

	// x := type{values} composite literal

	// slices like array but can change the length
	x := []int{1, 2, 3, 4, 5}

	//	x[5] = 6

	// add a new value to slice using append
	x = append(x, 6)

	fmt.Printf("%v\t%T\n", x, x)

	// a slice allows me to group values of same type

	// in slice you can go beyond the length
}
