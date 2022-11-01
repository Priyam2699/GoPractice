package main

import "fmt"

func main() {

	m := map[string]int{
		"James": 32,
		"Mark":  25,
		"C":     56,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Priyam"])

	v, k := m["Hello"]
	// v is the value and k is weather the key and value is present in map or not
	fmt.Println(v, k)

	if a, b := m["James"]; b {
		fmt.Println(a)
	}

}
