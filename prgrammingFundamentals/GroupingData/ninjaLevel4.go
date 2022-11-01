package main

import (
	"fmt"
)

func main() {

	map1 := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martini", "Women"},
		"moneypenny_miss": []string{"Literature", "computer science", "James Bond"},
		"no_dr":           []string{"Ice cream", "Evil", "Sunsets"},
	}

	fmt.Println(map1)

	map1["dua_priyam"] = []string{"Money", "Parents", "Friends"}

	for i, v1 := range map1 {

		fmt.Println("Record  ", i)

		for j, v2 := range v1 {

			fmt.Println("Key: ", j, "\tValue: ", v2, "\n")
		}

		fmt.Println()

	}
}
