package main

import "fmt"

func main() {
	foo()

	var s string = woo("Penny")

	fmt.Println(s)

	x, y := bloom("Ian", "Fleming")
	fmt.Println(x, y)

	sum := sum(3, 34, 5, 6, 4, 3, 2)
	fmt.Println(sum)

}

func sum(x ...int) int {
	fmt.Println("Adding up all the numbers")
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func foo() {
	fmt.Println("Foo")
}

func woo(s string) string {
	return fmt.Sprint("Hello dear ", s)
}

func bloom(x string, y string) (string, bool) {
	return fmt.Sprint(x, " ", y, ` says "Hello"`), false
}
