package main

import "fmt"

func main() {

	// s is a string
	s := `Hello "ABCD"`
	var str = `Hello "Iqbal"`

	fmt.Println(s)
	fmt.Println(str)

	bs := []byte(s)
	fmt.Printf("%T \n %v", bs, bs)

	for i, v := range s {
		fmt.Println(i, v)
	}

	var x = "世界"

	fmt.Println([]byte(x))
	fmt.Println([]rune(x))

}
