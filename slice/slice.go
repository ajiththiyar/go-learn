package main

import "fmt"

func main() {
	a := []int{2, 3, 4}
	fmt.Println(a)
	fmt.Println(a[1:])

	a = append(a, []int{34, 3}...)
	loopy(a)
	fmt.Println("Deleting 3 using slicing")
	// the semicolon side is actually expanded

	a = append(a[:1], a[2:]...)
	loopy(a)

	x := make([]int, 10, 12)

	fmt.Println(x)
	fmt.Println(`make uses array allocation method where 10 in this instance 
	denotes 10 inital zero values and 12 denotes starting allocation
	in memory which increases to 24 if it goes from 12 to 13`)
	x[0] = 3
	x[9] = 34
	x = append(x, 490)
	x = append(x, 490)
	loopy(x)
	x = append(x, 490)

	loopy(x)

}

func loopy(a []int) {
	fmt.Println("length", len(a))
	fmt.Println("capacity", cap(a))
	for _, value := range a {
		fmt.Println(value)
	}
}
