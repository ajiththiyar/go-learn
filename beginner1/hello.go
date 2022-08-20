package main

import "fmt"

var y int

func main() {
	value, e := fmt.Println("Hello, अश्वत्थामा")
	fmt.Println(value, e)
	fmt.Printf("%T\n", value)
	foo()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even", i)
			continue
		}
		fmt.Println("Odd", i)
	}
	y = 39
	fmt.Printf("Binary: %b\n", y)
	fmt.Printf("Hex: %x\n", y)
	fmt.Printf("Hex: %#x\n", y)

}

func foo() {
	fmt.Println("It works, FOO")
}
