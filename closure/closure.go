package main

import "fmt"

func main() {
	y := incrementor()
	fmt.Println(y())
	y()
	y()
	fmt.Println(y())
	x := incrementor()
	fmt.Println(x())
	x()
	x()
	fmt.Println(x())
}

// so the value inside scope is enclosed inside the func that is returned that is closure
func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
