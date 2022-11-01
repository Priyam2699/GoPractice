package main

import "fmt"

func main() {

	x := []string{"James", "Bond", "007"}
	y := []string{"Money", "Penny", "009"}

	fmt.Println(x)
	fmt.Println(y)

	multi := [][]string{x, y}

	fmt.Println(multi)

	// i is index and v is slices
	for i, v := range multi {

		fmt.Println("Record", i)

		for j, t := range v {
			fmt.Printf("\t Index Position: %v\tValue: %v\n", j, t)
		}
		fmt.Println()

	}

}
