package main

import "fmt"

var a int

// created my own type hotdog
type hotdog int

var b hotdog

func main() {

	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	b = 21
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// type conversion
	b = hotdog(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
