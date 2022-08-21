package main

import "fmt"

const (
	a = iota
	b
	c
	d
	e
)

func main() {

	x := 2
	y := x << 1
	fmt.Printf("%d    %b\n", x, x)
	fmt.Printf("%d    %b\n", y, y)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println("IOTA clever use")
	const (
		_ = iota
		l = 1 << (iota * 10)
		m = 1 << (iota * 10)
		n = 1 << (iota * 10)
	)
	fmt.Printf("%d    %b\n", l, l)
	fmt.Printf("%d    %b\n", m, m)
	fmt.Printf("%d    %b\n", n, n)

}
