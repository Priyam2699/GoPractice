package main

import "fmt"

func main() {

	s := 25
	fmt.Printf("%d\t%b\t%#x\n", s, s, s)
	t := s << 1
	fmt.Printf("%d\t%b\t%#x", t, t, t)
}
