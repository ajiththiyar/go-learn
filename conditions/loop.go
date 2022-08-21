package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println("")

	// create an array
	numbers := [5]int{11, 22, 33, 44, 55}

	// for loop with range
	for index := range numbers {
		fmt.Println(numbers[index])
	}
	fmt.Println("items")
	for _, item := range numbers {
		fmt.Println(item)
	}
	fmt.Println("ASCII")

	for i := 33; i < 122; i++ {
		var char = string(rune(i))
		fmt.Println(char)

	}

	fmt.Println("ASCII using int")
	for i := 33; i < 122; i++ {
		fmt.Printf("%#U\n", i)
	}

}
