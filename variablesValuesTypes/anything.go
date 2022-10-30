package main

import "fmt"

var y = 21
var z = "Shaken not stirred"
var x = true

// %T used to check the type of the variable
// is a static programming language
func main() {

	fmt.Println(y)
	fmt.Printf("%T \n", y)
	fmt.Printf("%T \n", z)
	fmt.Printf("%T \n", x)

}
