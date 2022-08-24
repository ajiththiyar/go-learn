package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 34
		close(c)
	}()

	v, ok := <-c
	if ok {

		fmt.Println(v, ok)
	}

	v, ok = <-c
	fmt.Println(v, ok)
}