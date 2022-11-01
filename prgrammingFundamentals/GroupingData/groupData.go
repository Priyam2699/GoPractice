package main

import "fmt"

func main() {

	// make function to assign capacity at last
	// ma(group Type,length,capacity)
	x := make([]int, 15, 15)

	fmt.Println(cap(x))
	fmt.Println(len(x))
	fmt.Println(x)

	x = append(x, 5)

	// if the size becomes greater than capacity then the capacity becomes double the original capacity
	fmt.Println(cap(x))
	fmt.Println(len(x))
	fmt.Println(x)

}
