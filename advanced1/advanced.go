package main

import (
	"fmt"
)

func doSomething(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int)

	go func() {
		ch <- doSomething(3)
	}()
	fmt.Println(<-ch)
}
