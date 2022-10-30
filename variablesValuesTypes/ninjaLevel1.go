package main

import "fmt"

var x int
var y string
var z bool

func main() {

	x = 42
	y = "James Bond"
	z = true
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Printf("%v %T\n", z, z)
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)

}
