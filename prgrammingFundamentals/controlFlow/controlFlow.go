package main

import "fmt"

func main() {

	for i := 1; i <= 3; i++ {

		fmt.Printf("Outer loop number %d\n", i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\tInner loop %d\n", j)

		}
		//fmt.Println(i)
	}

}
