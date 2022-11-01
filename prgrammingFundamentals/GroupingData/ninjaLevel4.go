package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 79, 50, 51}

	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, []int{56, 57, 58, 59, 60}...)

	fmt.Println(x)

}
