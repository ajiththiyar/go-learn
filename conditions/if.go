package main

import "fmt"

func main() {
	var a int = 34
	if a == 34 {
		fmt.Println("Works")
	}

	b := [3]int{3, 45, 6}
	fmt.Println(b)

	fmt.Println("Switch Statement")

	switch "bond" {
	case "bond":
		fmt.Println("bond")
	case "nope":
		fmt.Println("nope")
	default:
		fmt.Println("default")
	}

}
