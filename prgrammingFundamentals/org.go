package main

import "fmt"

var a int32
var b float32
var c int8 = 555 // as the range of int8 is -127 to 128

func main() {
	x := 42
	y := 42.1521
	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n", y)
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
}
