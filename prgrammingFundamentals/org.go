package main

import "fmt"

var x = 8

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%d\t\t%b\n", x, x)
	y := x >> 1
	fmt.Printf("%d\t\t%b\n", y, y)
	y = x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	kb := 1024
	mb := kb * 1024
	gb := mb * 1024

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

	// bit shifting using iota

	fmt.Printf("%d\t\t%b\n", KB, KB)
	fmt.Printf("%d\t\t%b\n", MB, MB)
	fmt.Printf("%d\t%b\n", GB, GB)

}
