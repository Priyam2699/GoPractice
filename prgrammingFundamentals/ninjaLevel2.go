package main

import "fmt"

const (
	a     = 42
	b int = 41
)

func main() {

	a := 42 == 42
	b := 42 >= 41
	c := 42 <= 41
	d := 42 != 41
	fmt.Println(a, b, c, d)

}
