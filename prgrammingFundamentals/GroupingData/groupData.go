package main

import "fmt"

func main() {

	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
	fmt.Println(x)

	var val int
	fmt.Scan(&val)

	for i, v := range x {

		if v == val {
			fmt.Println("Element found at position ", i)
			fmt.Println("Now deleting the element")
			x = append(x[:i], x[i+1:]...) // variadic parameter to take whole slice as an input
			fmt.Printf("The slice after deletion is %v", x)
			break
		}
	}

}
