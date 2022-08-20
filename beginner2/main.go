package main

import (
	"fmt"
)

// create your own type in golang using type keyword
type hotdog int

var b hotdog

func main() {
	a := 4
	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// conversion in go
	a = int(b)
	fmt.Println(a)
}
