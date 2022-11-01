package main

import "fmt"

func main() {

	jb := []string{"James", "Bond", "Vodka", "Martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Moneypenny", "At", "Your Service"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)

}
