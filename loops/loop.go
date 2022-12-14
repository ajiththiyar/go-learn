package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println("")

	// create an array
	numbers := []int{11, 22, 33, 44, 55}

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

	fmt.Println("Slicing")
	fmt.Println(numbers[2:])
	fmt.Println(numbers[2:3])

	morenumbers := []int{3, 4, 5, 6, 3}

	morenumbers = append(morenumbers, 53, 55, 34, 31)

	for _, value := range morenumbers {
		fmt.Println(value)
	}
	fmt.Println("dfgjkaldfgjfls;kdgjhkl;dfghjsdfghljk")
	numbers = append(numbers, morenumbers...)

	for _, value := range numbers {
		fmt.Println(value)
	}

}
