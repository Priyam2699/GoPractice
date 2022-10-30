package main

import "fmt"

var y = 911
var z = "Shaken not stirred"
var x = true

// %T used to check the type of the variable
// is a static programming language
func main() {

	fmt.Println(y)
	// %T gives type of the variable
	fmt.Printf("%T \n", y)
	// %b gives the binary value
	fmt.Printf("%b \n", y)
	fmt.Printf("%x \n", y)
	fmt.Printf("%#x \n", y)

	fmt.Printf("%T \n", z)
	fmt.Printf("%T \n", x)

	str := fmt.Sprintf("%#x \n", y)
	fmt.Println(len(str))
	fmt.Printf("%v", z) // %v gives real value

}
