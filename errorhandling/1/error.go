package main

import "fmt"

func main() {
	var name, food, age string
	fmt.Println("Name:")
	_, err := fmt.Scan(&name)

	if err != nil {
		panic(err)
	}

	fmt.Println("Food: ")
	_, err = fmt.Scan(&food)

	if err != nil {
		panic(err)
	}

	fmt.Println("Age: ")
	_, err = fmt.Scan(&age)

	if err != nil {
		panic(err)
	}

	fmt.Println(name, food, age)

}
