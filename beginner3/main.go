package main

import (
	"fmt"
)

// create your own type in golang using type keyword
type hotdog int

var b hotdog

func main() {
	a := "This is the beginning of String"
	fmt.Println(a)
	a = "अश्वत्थामा"
	bs := []byte(a)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(bs); i++ {
		fmt.Printf("%#U\n", bs[i])
	}

	for i, v := range bs {
		fmt.Println(i, v)
	}
}
