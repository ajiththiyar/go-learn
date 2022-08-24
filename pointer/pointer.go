package main

import "fmt"

func main() {
	a := 39
	fmt.Println(a)
	fmt.Println(&a)

	b := &a

	fmt.Println(b)
	fmt.Println(*b)
	a = 56
	fmt.Println(b)
	fmt.Println(*b)
}
