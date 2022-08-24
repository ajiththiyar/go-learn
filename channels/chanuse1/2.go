package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	const nums = 9
	wg.Add(nums * 2)
	for i := 0; i < 9; i++ {
		go foo(c, i)
	}
	for i := 0; i < 9; i++ {
		go bar(c)
	}
	wg.Wait()
	fmt.Println("Exiting")

}

func foo(c chan<- int, val int) {
	c <- val
	wg.Done()
}

func bar(c <-chan int) {
	fmt.Println(<-c)
	wg.Done()
}
