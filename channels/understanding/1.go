package main

import "fmt"

func main() {
	// buffer of 1
	c := make(chan int, 1)

	// r := make(chan <- int)   recieve only channel
	// s := make(<- chan int)   send only channel
	// c = s 					works
	// s = c 					does not work
	// c = r					works
	// r = c					does not work

	c <- 43

	fmt.Println(<-c)
}
