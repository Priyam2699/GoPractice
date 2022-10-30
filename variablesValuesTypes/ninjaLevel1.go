package main

import "fmt"

type mytype int

var x mytype
var y int

func main() {

	x = 42

	y = int(x)

	fmt.Printf("%v \n%T\n", y, y)

}
