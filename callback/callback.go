package main

import "fmt"

func main() {
	ii := []int{2, 34, 45, 56, 34, 2}
	fmt.Println(sum(ii...))

	fmt.Println("sum of even numbers happening using callback")
	fmt.Println(even(sum, ii...))
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func even(f func(i ...int) int, x ...int) int {
	var jj []int
	for _, v := range x {
		if v%2 == 0 {
			jj = append(jj, v)
		}
	}
	return f(jj...)
}
