package main

import "fmt"

func main() {

	var arr [5]int

	for i := 0; i < 5; i++ {

		arr[i] = i + 1

	}

	fmt.Println(arr)
	fmt.Println(len(arr))

}
