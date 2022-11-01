package main

import "fmt"

func main() {

	var favSport string
	fmt.Scan(&favSport)
	switch favSport {

	case "skiing":
		fmt.Println("Go to Blue Mountain")
		fallthrough
	case "swimming":
		fmt.Println("GO to Goa")
	case "surfing":
		fmt.Println("Go to Hawai")

	}

}
